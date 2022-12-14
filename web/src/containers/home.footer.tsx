import Image from "next/image";
import Logo from "@/assets/vercel.svg"

export const HomeFooter = () => {
    return (
        <footer className="container mx-auto p-24 flex flex-col gap-10">
            <Image src={Logo} alt={"logo"} className={"h-[50px] w-[100px]"}/>
            <p>Footer</p>
        </footer>
    )
}