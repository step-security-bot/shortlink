/* eslint-disable */

const webpack = require('webpack')
const withSourceMaps = require('@zeit/next-source-maps')

// ENVIRONMENT VARIABLE ================================================================================================
const isProd = process.env.NODE_ENV === 'production'
const isEnableSentry = process.env.SENTRY_ENABLE === 'true'
const API_URI = process.env.API_URI || 'http://localhost:7070'

console.info('API_URI', API_URI)

// You can choose which headers to add to the list
// after learning more below.
const securityHeaders = [
  {
    key: 'X-DNS-Prefetch-Control',
    value: 'on'
  },
  {
    key: 'X-XSS-Protection',
    value: '1; mode=block'
  },
  {
    key: 'X-Frame-Options',
    value: 'SAMEORIGIN'
  },
  {
    key: 'Permissions-Policy',
    value: 'camera=(), microphone=(), geolocation=(), interest-cohort=()'
  },
  {
    key: 'X-Content-Type-Options',
    value: 'nosniff'
  },
  {
    key: 'Referrer-Policy',
    value: 'origin-when-cross-origin'
  }
]

if (!isProd) {
  securityHeaders.push(
    {
      key: 'Strict-Transport-Security',
      value: 'max-age=63072000; includeSubDomains; preload'
    },
  )
}

const NEXT_CONFIG = {
  basePath: '/next',
  env: {
    NEXT_PUBLIC_API_URI: process.env.API_URI,
  },
  swcMinify: true,
  compiler: {
    // ssr and displayName are configured by default
    styledComponents: true,
  },
  images: {
    loader: "custom",
    // domains: ['images.unsplash.com', 'www.dropbox.com', 'www.tailwind-kit.com'],
    // formats: ['image/webp'],
  },
  webpack: (config, { isServer, buildId }) => {
    config.plugins.push(new webpack.DefinePlugin({}))

    return config
  },
  webpack5: true,
  trailingSlash: true,
  headers: () => {
    return [{
      // Apply these headers to all routes in your application.
      source: '/:path*',
      headers: securityHeaders,
    }]
  },
}

if (!isProd) {
  NEXT_CONFIG.rewrites = async function () {
    return {
      fallback: [
        // we need to define a no-op rewrite to trigger checking
        // all pages/static files before we attempt proxying
        {
          source: `/api`,
          destination: `${API_URI}/api`,
          basePath: false,
        },
        {
          source: `/api/:uri`,
          destination: `${API_URI}/api/:uri`,
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

let EXPORT_CONFIG = withSourceMaps(NEXT_CONFIG)

// Make sure adding Sentry options is the last code to run before exporting, to
// ensure that your source maps include changes from all other Webpack plugins
if (isEnableSentry) {
  // This file sets a custom webpack configuration to use your Next.js app
  // with Sentry.
  // https://nextjs.org/docs/api-reference/next.config.js/introduction
  // https://docs.sentry.io/platforms/javascript/guides/nextjs/
  const { withSentryConfig } = require('@sentry/nextjs')

  const SentryWebpackPluginOptions = {
    // Additional config options for the Sentry Webpack plugin. Keep in mind that
    // the following options are set automatically, and overriding them is not
    // recommended:
    //   release, url, org, project, authToken, configFile, stripPrefix,
    //   urlPrefix, include, ignore
    // For all available options, see:
    // https://github.com/getsentry/sentry-webpack-plugin#options.
  }

  EXPORT_CONFIG = withSentryConfig(EXPORT_CONFIG, SentryWebpackPluginOptions)
}

module.exports = EXPORT_CONFIG
