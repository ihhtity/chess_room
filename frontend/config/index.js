const path = require('path')

const config = {
  projectName: 'chess-room-frontend',
  date: '2024-01-01',
  designWidth: 750,
  deviceRatio: {
    640: 2.34 / 2,
    750: 1,
    828: 1.81 / 2
  },
  sourceRoot: 'src',
  outputRoot: 'dist',
  plugins: [],
  defineConstants: {},
  copy: {
    patterns: [],
    options: {}
  },
  framework: 'react',
  compiler: 'webpack5',
  cache: {
    enable: true
  },
  sass: {
    data: `@import "@nutui/nutui-react-taro/dist/styles/variables.scss";`,
    resource: [
      path.resolve(__dirname, '..', 'node_modules/@nutui/nutui-react-taro/dist/styles/variables.scss')
    ]
  },
  mini: {
    postcss: {
      pxtransform: {
        enable: true,
        config: {}
      },
      url: {
        enable: true,
        config: {
          limit: 10240
        }
      },
      cssModules: {
        enable: false,
        config: {
          namingPattern: 'module',
          generateScopedName: '[name]__[local]___[hash:base64:5]'
        }
      }
    },
    webpackChain: (chain, webpack) => {
      chain.resolve.alias.set('@', path.resolve(__dirname, '../src'))
      chain.module
        .rule('script')
        .use('babelLoader')
        .options({
          presets: [
            ['@babel/preset-env', { targets: { browsers: ['last 2 versions', 'ie >= 10'] } }],
            '@babel/preset-react',
            '@babel/preset-typescript'
          ],
          plugins: ['babel-plugin-transform-taroapi']
        })
    }
  },
  h5: {
    publicPath: '/',
    staticDirectory: 'static',
    postcss: {
      autoprefixer: {
        enable: true,
        config: {}
      },
      cssModules: {
        enable: false,
        config: {
          namingPattern: 'module',
          generateScopedName: '[name]__[local]___[hash:base64:5]'
        }
      }
    },
    webpackChain: (chain, webpack) => {
      chain.resolve.alias.set('@', path.resolve(__dirname, '../src'))
      chain.plugins.delete('progress')
      chain.plugin('progress').use(webpack.ProgressPlugin, [{}])
      chain.module
        .rule('script')
        .use('babelLoader')
        .options({
          presets: [
            ['@babel/preset-env', { targets: { browsers: ['last 2 versions', 'ie >= 10'] } }],
            '@babel/preset-react',
            '@babel/preset-typescript'
          ],
          plugins: ['babel-plugin-transform-taroapi']
        })
    }
  }
}

module.exports = function (merge) {
  if (process.env.NODE_ENV === 'development') {
    return merge({}, config, require('./dev'))
  }
  return merge({}, config, require('./prod'))
}