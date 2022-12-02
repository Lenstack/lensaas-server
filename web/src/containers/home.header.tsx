import {Navigation} from "@/components";
import {HOME_LINKS} from "@/constants";

export const HomeHeader = () => {
    return (
        <header className="mx-auto py-10">
            <Navigation className="relative flex justify-around mx-auto max-w-7xl font-thin">
                <Navigation.Container className={"flex items-center gap-5"}>
                    {
                        HOME_LINKS.map(({id, title, url}) => (
                            <Navigation.Item key={id} href={url}>{title}</Navigation.Item>
                        ))
                    }
                </Navigation.Container>
                <Navigation.Container className={"flex items-center gap-5"}>
                    <Navigation.Item href="/dashboard">Sign In</Navigation.Item>
                    <Navigation.Item href="/">Sign Up</Navigation.Item>
                </Navigation.Container>
            </Navigation>
        </header>
    )
}