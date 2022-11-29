import Head from "next/head";
import {Landing} from "@/containers";

export default function Home() {
    return (
        <>
            <Head>
                <title>Lensaas</title>
                <meta charSet="UTF-8"/>
                <meta name="description" content="Template For Build Saas"/>
                <meta name="author" content="leonardo ospina"/>
                <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
                <link rel="icon" href="/favicon.ico"/>
            </Head>
            <Landing/>
        </>
    )
}