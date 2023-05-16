import Document, { Head, Html, Main, NextScript } from 'next/document'
import createEmotionServer from '@emotion/server/create-instance'
import { getInitColorSchemeScript } from '@mui/material/styles'
// @ts-ignore
import { createEmotionCache, lightTheme } from '@shortlink-org/ui-kit'
import PropTypes from 'prop-types'

export default function MyDocument(props: any) {
  const { emotionStyleTags } = props

  return (
    <Html lang="en">
      <Head>
        {/* PWA primary color */}
        <meta name="theme-color" content={lightTheme.palette.primary.main} />
        <link rel="manifest" href="/manifest.json" />

        <link rel="shortcut icon" href="/favicon.ico" />
        <link
          rel="stylesheet"
          href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
        />
        <script
          src="https://challenges.cloudflare.com/turnstile/v0/api.js"
          async
          defer
        />
        {/* Inject MUI styles first to match with to prepend: true configuration. */}
        {emotionStyleTags}
      </Head>
      <body className="bg-white text-black dark:bg-gray-800 dark:text-white">
        {getInitColorSchemeScript()}
        <Main />
        <NextScript />
      </body>
    </Html>
  )
}

// `getInitialProps` belongs to `_document` (instead of `_app`),
// it's compatible with static-site generation (SSG).
// @ts-ignore
MyDocument.getInitialProps = async (ctx) => {
  // Resolution order
  //
  // On the server:
  // 1. app.getInitialProps
  // 2. page.getInitialProps
  // 3. document.getInitialProps
  // 4. app.render
  // 5. page.render
  // 6. document.render
  //
  // On the server with error:
  // 1. document.getInitialProps
  // 2. app.render
  // 3. page.render
  // 4. document.render
  //
  // On the client
  // 1. app.getInitialProps
  // 2. page.getInitialProps
  // 3. app.render
  // 4. page.render

  const originalRenderPage = ctx.renderPage

  // You can consider sharing the same Emotion cache between all the SSR requests to speed up performance.
  // However, be aware that it can have global side effects.
  const cache = createEmotionCache()
  const { extractCriticalToChunks } = createEmotionServer(cache)

  ctx.renderPage = () =>
    originalRenderPage({
      enhanceApp: (
        App: any, // @ts-ignore
      ) => // @ts-ignore
        function EnhanceApp(props) {
          return <App emotionCache={cache} {...props} />
        },
    })

  const initialProps = await Document.getInitialProps(ctx)
  // This is important. It prevents Emotion to render invalid HTML.
  // See https://github.com/mui/material-ui/issues/26561#issuecomment-855286153
  const emotionStyles = extractCriticalToChunks(initialProps.html)
  const emotionStyleTags = emotionStyles.styles.map((style) => (
    <style
      data-emotion={`${style.key} ${style.ids.join(' ')}`}
      key={style.key}
      // eslint-disable-next-line react/no-danger
      dangerouslySetInnerHTML={{ __html: style.css }}
    />
  ))

  return {
    ...initialProps,
    emotionStyleTags,
  }
}

MyDocument.propTypes = {
  // eslint-disable-next-line react/forbid-prop-types
  emotionStyleTags: PropTypes.array.isRequired,
}
