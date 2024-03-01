/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./templates/**/*.{html,js,templ,go}",
      "./templates/common/**/*.{html,js,templ,go}",
    ],
    theme: {
      extend: {},
      fontFamily: {
        sans: ["Quicksand"],
        body: ['Noto Sans']
      },
    },
    plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
  };
  