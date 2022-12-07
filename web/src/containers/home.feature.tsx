import Image from "next/image";
import Dashboard from "@/assets/dashboard.jpg";

export const HomeFeature = () => {
    return (
        <section className="container mx-auto py-5">
            <section className={"flex flex-col gap-10 text-center"}>
                <h2 className={"text-5xl font-bold text-center"}>How works</h2>
                <p className={"text-2xl text-gray-700"}>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi pariatur reprehenderit!</p>
            </section>
            <section className={"p-28"}>
                <Image src={Dashboard} alt={"logo"} className={" object-fill rounded-2xl"}/>
            </section>
        </section>
    )
}