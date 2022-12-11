import {HomeContact, HomeFeature, HomeFooter, HomeHeader, HomeHero, HomePricing} from "@/containers";

export default function Home() {
    return (
        <div className={"min-h-screen overflow-hidden"}>
            <HomeHeader/>
            <HomeHero/>
            <HomeFeature/>
            <HomePricing/>
            <HomeContact/>
            <HomeFooter/>
        </div>
    )
}