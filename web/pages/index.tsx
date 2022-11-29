import {Navigation} from "../components";

export default function Home() {
    return (
        <>
            <Navigation>
                <Navigation.Item href="/">Home</Navigation.Item>
                <Navigation.Item href="/dashboard">Dashboard</Navigation.Item>
                <Navigation.Item href="/sign_up">Sign Up</Navigation.Item>
                <Navigation.Item href="/sign_in">Sign In</Navigation.Item>
            </Navigation>
        </>
    )
}
