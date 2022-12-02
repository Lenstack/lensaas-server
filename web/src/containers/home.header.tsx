import {Navigation} from "@/components";

export const HomeHeader = () => {
    return (
        <header className="text-gray-900 font-normal bg-gray-100">
            <Navigation className="m-3 flex justify-around p-7">
                <Navigation.Container className={"flex justify-between gap-5"}>
                    <Navigation.Item href="/">Logo</Navigation.Item>
                    <Navigation.Item href="/">Features</Navigation.Item>
                    <Navigation.Item href="/">Pricing</Navigation.Item>
                    <Navigation.Item href="/">Contact</Navigation.Item>
                </Navigation.Container>
                <Navigation.Container className={"flex justify-between gap-5"}>
                    <Navigation.Item href="/dashboard">Dashboard</Navigation.Item>
                    <Navigation.Item href="/">Sign In</Navigation.Item>
                    <Navigation.Item href="/">Sign Up</Navigation.Item>
                </Navigation.Container>
            </Navigation>
        </header>
    )
}