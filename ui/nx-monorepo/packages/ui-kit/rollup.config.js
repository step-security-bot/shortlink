import { nodeResolve } from '@rollup/plugin-node-resolve'
import commonjs from '@rollup/plugin-commonjs'
import ts from '@rollup/plugin-typescript'
import typescript from 'typescript'
import peerDepsExternal from 'rollup-plugin-peer-deps-external'
import terser from '@rollup/plugin-terser'
import postcss from 'rollup-plugin-postcss'
// import dts from 'rollup-plugin-dts'
import filesize from 'rollup-plugin-filesize'
import svgr from '@svgr/rollup'

// eslint-disable-next-line @typescript-eslint/no-var-requires
const packageJson = require('./package.json')

const globals = {
  react: 'React',
  'react-dom': 'ReactDOM',
  classnames: 'classNames',
  'prop-types': 'PropTypes',
  '@emotion/styled': 'emStyled',
  '@emotion/react': 'react',
  'react/jsx-runtime': 'jsxRuntime',
}

export default [
  {
    input: 'src/index.ts',
    output: [
      {
        name: packageJson.name,
        file: packageJson.main,
        format: 'umd',
        sourcemap: true,
        globals,
      },
      {
        file: packageJson.module,
        format: 'esm',
        sourcemap: true,
      },
    ],
    plugins: [
      peerDepsExternal(), // Preferably set as first plugin.
      svgr(),
      postcss({
        extract: 'index.css',
        autoModules: true,
        include: 'src/**/*.css',
        extensions: ['.css'],
        plugins: [],
      }),
      nodeResolve(),
      ts({
        typescript,
        tsconfig: './tsconfig.json',
        noEmitOnError: false,
        declaration: true,
      }),
      commonjs({
        include: 'node_modules/**',
      }),
      filesize(),
      terser(),
    ],
    external: [
      'react',
      'react-is',
      'react-dom',
      'prop-types',
      'styled-components',
      '@mui/material',
      '@emotion/react',
    ],
  },
  // {
  //   name: packageJson.name,
  //   input: "dist/esm/types/index.d.ts",
  //   output: [{ file: "dist/index.d.ts", format: "es" }],
  //   plugins: [
  //     postcss({
  //       extract: 'styles.css', // this will generate a specific file not being used, but we need this part of code
  //       autoModules: true,
  //       include: '**/*.css',
  //       extensions: ['.css'],
  //       plugins: [],
  //     }),
  //     dts.default({
  //       compilerOptions: {
  //         baseUrl: 'src',
  //       },
  //     }),
  //   ],
  //   external: ['react', 'react-is', 'react-dom', 'prop-types', 'styled-components', "@mui/material", "@emotion/react"],
  // },
]
