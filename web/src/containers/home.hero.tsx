export const HomeHero = () => {
    return (
        <main className={"container mx-auto p-24 flex flex-col gap-20"}>
            <section className={"flex flex-col text-center"}>
                <h1 className="text-8xl font-bold my-10 text-transparent bg-clip-text bg-gradient-to-l from-orange-400 via-cyan-300 to-pink-800">
                    The Fastest Way To Build Your SAAS With Microservices
                </h1>
                <p className={"text-2xl text-gray-700"}>
                    Boost your speediness in building your saas with microservices
                </p>
            </section>
            <section className={"flex align-middle justify-center gap-5"}>
                <button
                    className={"py-4 rounded-full px-7 bg-zinc-800 text-white border-2 hover:border-zinc-200"}>Demo
                </button>
                <button>Documentation</button>
            </section>
        </main>
    )
}