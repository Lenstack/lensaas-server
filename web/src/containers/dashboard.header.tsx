import {Avatar, Navigation} from "@/components";

export const DashboardHeader = () => {
    return (
        <header className="col-start-3 col-end-12 row-start-1 row-end-2 p-10 flex justify-between items-center p-5">
            <h1 className="text-2xl font-bold">
                Module
            </h1>
            <Navigation className="flex justify-center items-center gap-10">
                <Navigation.Container className="flex gap-5">
                    <Navigation.Item href="/dashboard">Dashboard</Navigation.Item>
                    <Navigation.Item href="/">Home</Navigation.Item>
                </Navigation.Container>
                <Navigation.Container>
                    <Avatar/>
                </Navigation.Container>
            </Navigation>
        </header>
    )
}