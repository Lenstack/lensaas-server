import {HomeContact, HomeFeature, HomeFooter, HomeHeader, HomeHero, HomePricing} from "@/containers";

export default function Home() {
    return (
        <div className={"bg-gradient-to-b from-gray-100 to-gray-300 min-h-screen overflow-hidden"}>
            <HomeHeader/>
            <HomeHero/>
            <HomeFeature/>
            <HomePricing/>
            <HomeContact/>
            <HomeFooter/>
        </div>
    )
}