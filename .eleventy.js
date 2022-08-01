const navigation = require('@11ty/eleventy-navigation')
const CleanCSS = require("clean-css");
const htmlmin = require("html-minifier");

module.exports = function (eleventyConfig) {
    // Plugins
    eleventyConfig.addPlugin(navigation)

    // Layout Aliases
    eleventyConfig.addLayoutAlias('base', 'layouts/base.njk')

    // Passthroughs
    eleventyConfig.addPassthroughCopy({
        'src/_includes/assets/css/style.css': './style.css'
    });
    eleventyConfig.addPassthroughCopy({
        'src/_includes/assets/img': './img'
    });

    // Filters
    eleventyConfig.addFilter("cssmin", function (code) {
        return new CleanCSS({}).minify(code).styles;
    });

    // Transforms
    eleventyConfig.addTransform("htmlmin", function (content) {
        if (this.outputPath && this.outputPath.endsWith(".html")) {
            let minified = htmlmin.minify(content, {
                useShortDoctype: true,
                removeComments: true,
                collapseWhitespace: true
            });
            return minified;
        }
        return content;
    });

    return {
        dir: {
            input: "src",
            output: "dist",
        },
        dataTemplateEngine: "njk"
    }
}