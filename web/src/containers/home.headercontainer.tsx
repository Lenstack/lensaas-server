import {Navigation} from "@/components";

export const HomeHeaderContainer = () => {
    return (
        <header className="">
            <Navigation>
                <Navigation.Item href="/">Home</Navigation.Item>
                <Navigation.Item href="/dashboard">Dashboard</Navigation.Item>
                <Navigation.Container>
                    <button className="">
                        Sign In
                    </button>
                    <button className="">
                        Sign Up
                    </button>
                </Navigation.Container>
            </Navigation>
        </header>
    )
}