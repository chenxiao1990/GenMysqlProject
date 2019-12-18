module.exports = {
  publicPath: '/vue/',
  outputDir: undefined,
  assetsDir: undefined,
  runtimeCompiler: undefined,
  productionSourceMap: false,
  parallel: undefined,
  css: undefined,
  chainWebpack: config => {
    config.plugins.delete('prefetch')
  }
}