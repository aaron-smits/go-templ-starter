/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      'view/**/*.templ',
    ],
    darkMode: 'class',
    theme: {
    // //   extend: {
    // //     fontFamily: {
    // //       mono: ['Courier Prime', 'monospace'],
    // //     }
    // },
    },

    plugins: [
      require('tailwindcss'),
    ],
    corePlugins: {
      preflight: true,
    }
  }