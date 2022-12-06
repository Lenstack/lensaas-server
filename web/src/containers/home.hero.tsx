export const HomeHero = () => {
    return (
        <main className={"container mx-auto p-64 flex flex-col gap-20"}>
            <section className={"flex flex-col text-center"}>
                <h1 className="text-8xl font-bold my-10 ">
                    The Fastest Way To Build Your <span
                    className={"text-transparent bg-clip-text bg-gradient-to-l from-orange-400 via-cyan-300 to-pink-800"}>SAAS</span> With <span className={"text-transparent bg-clip-text bg-gradient-to-r from-green-300 via-blue-500 to-purple-600"}>Microservices</span>
                </h1>
                <p className={"text-2xl text-gray-700"}>
                    We provide you with the best tools to build your saas with microservices.
                </p>
            </section>
            <section className={"flex align-middle justify-center gap-5"}>
                <button
                    className={"py-5 px-10 rounded-full bg-zinc-800 text-white border-2 hover:border-zinc-200"}>Demo
                </button>
                <button>Documentation</button>
            </section>
        </main>
    )
}