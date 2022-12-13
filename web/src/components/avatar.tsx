import Image from "next/image";
import AvatarSVG from "@/assets/avatars/avatar3.svg";

export const Avatar = () => {
    return (
        <div className="flex flex-col items-center gap-2">
            <div className="w-12 h-12 rounded-full bg-[#F6F6F6] dark:bg-zinc-800 dark:text-white">
                <Image src={AvatarSVG} alt={"avatar"} className="w-full h-full rounded-2xl"/>
            </div>
        </div>
    )
}