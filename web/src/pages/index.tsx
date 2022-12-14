import {HomeContact, HomeFeature, HomeFooter, HomeFq, HomeHeader, HomeHero, HomePricing} from "@/containers";

export default function Home() {
    return (
        <div className="min-h-screen overflow-hidden">
            <HomeHeader/>
            <HomeHero/>
            <HomeFeature/>
            <HomePricing/>
            <HomeFq/>
            <HomeContact/>
            <HomeFooter/>
        </div>
    )
}