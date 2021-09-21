const colors = require('tailwindcss/colors')

module.exports = {
    purge: [],
    darkMode: false, // or 'media' or 'class'
    theme: {
        minHeight: {
            '0': '0',
            '1/4': '25%',
            '1/2': '50%',
            '3/4': '75%',
            'full': '100%',
        },
        minWidth: {
            '0': '0',
            '1/4': '25%',
            '1/2': '50%',
            '3/4': '75%',
            'full': '100%',
        },
        colors: {
            orange: {
                DEFAULT: '#FF7900'
            },
            gray: {
                DEFAULT: '#8F8F8F',
                light: '#D6D6D6',
                dark: '#595959'
            },
            grey: colors.gray,
            indigo: colors.indigo,
            blue: {
                DEFAULT: '#4BB4E6'
            },
            green: {
                DEFAULT: '#50BE87'
            },
            pink: {
                DEFAULT: '#FFB4E6'
            },
            purple: {
                DEFAULT: '#A885D8'
            },
            yellow: {
                DEFAULT: '#FFD200'
            },
            red: {
                DEFAULT: '#e11d48'
            },
            black: colors.black,
            white: colors.white,
        },
        extend: {},
    },
    variants: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
    ],
}
