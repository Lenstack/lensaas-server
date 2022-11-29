import {Navigation} from "@/components";

export const Landing = () => {
    return (
        <>
            <header>
                <Navigation>
                    <Navigation.Item href="/">Home</Navigation.Item>
                    <Navigation.Item href="/dashboard">Dashboard</Navigation.Item>
                    <Navigation.Container>
                        <Navigation.Item href="/sign_in">Sign In</Navigation.Item>
                        <Navigation.Item href="/sign_up">Sign Up</Navigation.Item>
                    </Navigation.Container>
                </Navigation>
            </header>
            <main>
                <h1> Landing Main </h1>
            </main>
        </>
    )
}