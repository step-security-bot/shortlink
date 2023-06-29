/* eslint-disable */
const { withSentryConfig } = require('@sentry/nextjs')
const withPWA = require('@ducanh2912/next-pwa').default({
  dest: 'public',
  maximumFileSizeToCacheInBytes: 5000000,
})

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'
const isEnableSentry = process.env.SENTRY_ENABLE === 'true'
const API_URI = process.env.API_URI || 'http://localhost:7070'

console.info('API_URI', API_URI)

// You can choose which headers to add to the list
// after the learning more below.
const securityHeaders = [
  {
    key: 'X-DNS-Prefetch-Control',
    value: 'on',
  },
  {
    key: 'X-XSS-Protection',
    value: '1; mode=block',
  },
  {
    key: 'X-Frame-Options',
    value: 'SAMEORIGIN',
  },
  {
    key: 'Permissions-Policy',
    value: 'camera=(), microphone=(), geolocation=(), interest-cohort=()',
  },
  {
    key: 'X-Content-Type-Options',
    value: 'nosniff',
  },
  {
    key: 'Referrer-Policy',
    value: 'origin-when-cross-origin',
  },
]

if (!isProd) {
  securityHeaders.push({
    key: 'Strict-Transport-Security',
    value: 'max-age=63072000; includeSubDomains; preload',
  })
}

let NEXT_CONFIG = {
  basePath: '/next',
  generateEtags: true,
  env: {
    // ShortLink API
    NEXT_PUBLIC_SERVICE_NAME: 'shortlink-next',
    NEXT_PUBLIC_API_URI: process.env.API_URI,

    // Sentry
    NEXT_PUBLIC_SENTRY_DSN: process.env.SENTRY_DSN,

    // Firebase
    NEXT_PUBLIC_FIREBASE_VAPID_KEY: process.env.FIREBASE_VAPID_KEY,
    NEXT_PUBLIC_FIREBASE_API_KEY: process.env.FIREBASE_API_KEY,
    NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN: process.env.FIREBASE_AUTH_DOMAIN,
    NEXT_PUBLIC_FIREBASE_PROJECT_ID: process.env.FIREBASE_PROJECT_ID,
    NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET: process.env.FIREBASE_STORAGE_BUCKET,
    NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID:
      process.env.FIREBASE_MESSAGING_SENDER_ID,
    NEXT_PUBLIC_FIREBASE_APP_ID: process.env.FIREBASE_APP_ID,
    NEXT_PUBLIC_FIREBASE_MEASUREMENT_ID: process.env.FIREBASE_MEASUREMENT_ID,
  },
  swcMinify: true,
  productionBrowserSourceMaps: true,
  transpilePackages: ['@shortlink-org/ui-kit'],
  compiler: {
    // ssr and displayName are configured by default
    styledComponents: true,
  },
  images: {
    loader: 'custom',
    domains: ['images.unsplash.com'],
    formats: ['image/avif', 'image/webp'],
    remotePatterns: [
      {
        // The `src` property hostname must end with `.example.com`,
        // otherwise the API will respond with 400 Bad Request.
        protocol: 'https',
        hostname: 'images.unsplash.com',
      },
    ],
    dangerouslyAllowSVG: false,
    contentDispositionType: 'attachment',
    contentSecurityPolicy: "default-src 'self'; script-src 'none'; sandbox;",
  },
  trailingSlash: false,
  headers: () => {
    return [
      {
        // Apply these headers to all routes in your application.
        source: '/:path*',
        headers: [
          ...securityHeaders,
          {
            key: 'Cache-Control',
            value:
              'public, max-age=14400, s-maxage=14400, stale-while-revalidate=86400',
          },
        ],
      },
    ]
  },
  webpack: (config, { isServer, buildId }) => {
    config.module.rules.push({
      test: /\.svg$/i,
      issuer: /\.[jt]sx?$/,
      use: ['@svgr/webpack'],
    })

    return config
  },
  experimental: {
    forceSwcTransforms: true,
    turbo: {
      loaders: {
        '.md': [
          {
            loader: '@mdx-js/loader',
            options: {
              format: 'md',
            },
          },
        ],
        // Option-less format
        '.mdx': '@mdx-js/loader',
      },
    },
    swcTraceProfiling: true,
  },
  sentry: {
    hideSourceMaps: true,
  },
}

if (!isProd) {
  NEXT_CONFIG.rewrites = async function () {
    return {
      fallback: [
        // we need to define a no-op rewrite to trigger checking
        // all pages/static files before we attempt proxying
        {
          source: `/api/auth/:uri*`,
          destination: `http://127.0.0.1:4433/:uri*`,
          basePath: false,
        },
        {
          source: `/api/:uri*`,
          destination: `http://127.0.0.1:7070/api/:uri*`,
          basePath: false,
        },
        {
          source: `/s`,
          destination: `${API_URI}/s`,
          basePath: false,
        },
        {
          source: `/s/:uri`,
          destination: `${API_URI}/s/:uri`,
          basePath: false,
        },
      ],
    }
  }
}

// Make sure adding Sentry options is the last code to run before exporting, to
// ensure that your source maps include changes from all other Webpack plugins
if (isEnableSentry) {
  const SentryWebpackPluginOptions = {
    // For all available options, see:
    // https://github.com/getsentry/sentry-webpack-plugin#options

    // Suppresses source map uploading logs during build
    silent: true,

    org: 'batazor',
    project: 'shortlink-next',
  }

  NEXT_CONFIG = withSentryConfig(NEXT_CONFIG, SentryWebpackPluginOptions)
}

module.exports = withPWA(NEXT_CONFIG)
