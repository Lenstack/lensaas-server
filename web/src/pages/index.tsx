import {HomeContact, HomeFeature, HomeFooter, HomeHeader, HomeHero, HomePricing} from "@/containers";

export default function Home() {
    return (
        <div className={"bg-[#FFFFFF] h-screen"}>
            <HomeHeader/>
            <HomeHero/>
            <HomeFeature/>
            <HomePricing/>
            <HomeContact/>
            <HomeFooter/>
        </div>
    )
}