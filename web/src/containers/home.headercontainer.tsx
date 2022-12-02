import {Navigation} from "@/components";

export const HomeHeaderContainer = () => {
    return (
        <header className="text-gray-900 font-normal">
            <Navigation className="m-3 flex justify-around p-4">
                <Navigation.Container>
                    <Navigation.Item href="/">Home</Navigation.Item>
                    <Navigation.Item href="/dashboard">Dashboard</Navigation.Item>
                    <Navigation.Item href="/about">About</Navigation.Item>
                </Navigation.Container>
                <Navigation.Container>
                    <Navigation.Button>
                        Sign In
                    </Navigation.Button>
                    <Navigation.Button className="bg-stone-800 text-white p-2 rounded">
                        Sign Up
                    </Navigation.Button>
                </Navigation.Container>
            </Navigation>
        </header>
    )
}