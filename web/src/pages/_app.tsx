import '../styles/globals.css'
import type {AppProps} from 'next/app'
import {ThemeProvider} from "next-themes";
import Head from "next/head";

export default function App({Component, pageProps}: AppProps) {
    return (
        <ThemeProvider attribute="class">
            <Head>
                <meta charSet="utf-8"/>
                <meta name="viewport" content="width=device-width, initial-scale=1"/>
                <meta name="description" content="Lensaas With Microservices"/>
                <meta name="author" content="Lenstack"/>
                <meta name="keywords" content="keywords"/>
                <meta name="theme-color" content="#EEF0F2"/>
                <link rel="manifest" href="/manifest.json" />
                <title>Lensaas</title>
            </Head>
            <Component {...pageProps} />
        </ThemeProvider>
    )
}
