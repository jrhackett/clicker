const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const path = require('path')

const mode = process.env.NODE_ENV || 'development'
const prod = mode === 'production'

module.exports = {
    devServer: {
        historyApiFallback: true,
    },
    devtool: prod ? false : 'source-map',
	entry: {
		bundle: ['./main.js']
    },
    mode,
    module: {
		rules: [
			{
				test: /\.svelte$/,
				use: {
					loader: 'svelte-loader',
					options: {
						dev: prod ? false : true,
						emitCss: true,
						hotReload: true
					}
				}
			},
			{
				test: /\.css$/,
				use: [
					/**
					 * MiniCssExtractPlugin doesn't support HMR.
					 * For developing, use 'style-loader' instead.
					 * */
					prod ? MiniCssExtractPlugin.loader : 'style-loader',
					'css-loader'
				]
			}
		]
	},
    output: {
		path: __dirname + '/public',
		filename: '[name].js',
        chunkFilename: '[name].[id].js',
        publicPath: '/'
    },
    plugins: [
		new MiniCssExtractPlugin({ filename: '[name].css' })
	],
	resolve: {
		alias: {
			svelte: path.resolve('node_modules', 'svelte')
		},
		extensions: ['.mjs', '.js', '.svelte'],
        mainFields: ['svelte', 'browser', 'module', 'main'],
        modules: [
            path.resolve('./'),
            path.resolve('./node_modules')
        ]
	}
}
