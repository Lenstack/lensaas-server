import {Navigation} from "@/components";

export const DashboardAside = () => {
    return (
        <aside className="bg-amber-100">
            <Navigation>
                <Navigation.Container className={"flex flex-col gap-5"}>
                    <Navigation.Item href="/">Home</Navigation.Item>
                    <Navigation.Item href="/dashboard">Dashboard</Navigation.Item>
                </Navigation.Container>
            </Navigation>
        </aside>
    )
}