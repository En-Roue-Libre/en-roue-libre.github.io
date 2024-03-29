const navigation = require('@11ty/eleventy-navigation')
const CleanCSS = require("clean-css");
const htmlmin = require("html-minifier");
const fg = require('fast-glob');
const galleryImages = fg.globSync('**/img/gallery/*', { onlyFiles: true, cwd: 'src/_includes/assets' });
const memeImages = fg.globSync('**/img/memes/*', { onlyFiles: true, cwd: 'src/_includes/assets' });

module.exports = function (eleventyConfig) {
    // Collections
    eleventyConfig.addCollection('gallery', function () {
        return galleryImages;
    });
    eleventyConfig.addCollection('memes', function () {
        return memeImages;
    });

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
    eleventyConfig.addPassthroughCopy({
        'src/_includes/assets/favicon': '/'
    });

    // Filters
    eleventyConfig.addFilter("cssmin", function (code) {
        return new CleanCSS({}).minify(code).styles;
    });
    eleventyConfig.addFilter("toLocaleDateString", function (timestamp) {
        return new Date(timestamp * 1000).toLocaleDateString('fr-FR', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' });
    });
    eleventyConfig.addFilter("isPast", function (timestamp) {
        return new Date(timestamp * 1000) < new Date();
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