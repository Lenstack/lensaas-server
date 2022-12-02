import {HomeContact, HomeFeature, HomeFooter, HomeHeader, HomeHero, HomePricing} from "@/containers";

export default function Home() {
    return (
        <div className={"w-full h-full overflow-hidden"}>
            <HomeHeader/>
            <HomeHero/>
            <HomeFeature/>
            <HomePricing/>
            <HomeContact/>
            <HomeFooter/>
        </div>
    )
}