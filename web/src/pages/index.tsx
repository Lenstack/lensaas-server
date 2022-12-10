import {HomeContact, HomeFeature, HomeFooter, HomeHeader, HomeHero, HomePricing} from "@/containers";

export default function Home() {
    return (
        <div className={"min-h-screen overflow-hidden bg-gradient-to-b from-gray-100 to-gray-300" +
            "dark:bg-gradient-to-br dark:from-zinc-900 dark:via-zinc-900 dark:to-neutral-900"}>
            <HomeHeader/>
            <HomeHero/>
            <HomeFeature/>
            <HomePricing/>
            <HomeContact/>
            <HomeFooter/>
        </div>
    )
}