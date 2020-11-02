package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path"
	"syscall"

	csi_driver "github.com/batazor/shortlink/pkg/csi"
	"github.com/batazor/shortlink/pkg/csi/di"
)

func main() {
	// Create a new context
	ctx := context.Background()

	// Init a new service
	s, cleanup, err := di.InitializeSCIDriver(ctx)
	if err != nil { // TODO: use as helpers
		if r, ok := err.(*net.OpError); ok {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT enviroment", r.Addr.String()))
		}

		panic(err)
	}

	// TODO: Use cobra
	var (
		endpoint          = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
		nodeID            = flag.String("nodeid", "", "node id")
		ephemeral         = flag.Bool("ephemeral", false, "publish volumes in ephemeral mode even if kubelet did not ask for it (only needed for Kubernetes 1.15)")
		maxVolumesPerNode = flag.Int64("maxvolumespernode", 0, "limit of volumes per node")
		showVersion       = flag.Bool("version", false, "Show version.")
		// Set by the build process
		version = ""
	)

	flag.Parse()

	if *showVersion {
		baseName := path.Base(os.Args[0])
		fmt.Println(baseName, version)
		return
	}

	if *ephemeral {
		fmt.Fprintln(os.Stderr, "Deprecation warning: The ephemeral flag is deprecated and should only be used when deploying on Kubernetes 1.15. It will be removed in the future.")
	}

	// Run CSI Driver
	driver, err := csi_driver.NewDriver(s.Log, "shrts.csi.k8s.io", *nodeID, *endpoint, *ephemeral, *maxVolumesPerNode, version)
	if err != nil {
		s.Log.Fatal(fmt.Sprintf("Failed to initialize driver: %s", err.Error()))
	}
	if err := driver.Run(ctx); err != nil {
		s.Log.Fatal(err.Error())
	}

	s.Log.Info("success run CSI plugin")

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	cleanup()
}
