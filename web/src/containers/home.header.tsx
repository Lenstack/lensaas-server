import Image from "next/image";
import {Navigation} from "@/components";
import {HOME_LINKS} from "@/constants";
import Logo from "@/assets/vercel.svg"

export const HomeHeader = () => {
    return (
        <header className={"container"}>
            <Navigation className="flex justify-around align-middle pt-5 pb-5 fixed w-full bg-[#F0F1F3] bg-opacity-80">
                <Navigation.Container className={"flex items-center gap-5"}>
                    <Navigation.Item href="/">
                        <Image src={Logo} alt={"logo"} className={"h-[50px] w-[100px] object-fill"}/>
                    </Navigation.Item>
                </Navigation.Container>
                <Navigation.Container className={"flex items-center gap-5"}>
                    {
                        HOME_LINKS.map(({id, title, url}) => (
                            <Navigation.Item key={id} href={url}>{title}</Navigation.Item>
                        ))
                    }
                </Navigation.Container>
                <Navigation.Container className={"flex items-center gap-5"}>
                    <Navigation.Item href="/dashboard">Sign In</Navigation.Item>
                    <Navigation.Item href="/"
                                     className={"py-2 px-4 rounded-full border-zinc-800 border hover:bg-transparent hover:bg-zinc-800 hover:border-zinc-800 hover:text-white"}>Sign
                        Up</Navigation.Item>
                </Navigation.Container>
            </Navigation>
        </header>
    )
}