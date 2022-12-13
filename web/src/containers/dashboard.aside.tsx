import {Navigation, ThemeToggle} from "@/components";
import {DASHBOARD_LINKS} from "@/constants";

export const DashboardAside = () => {
    return (
        <aside
            className="bg-[#F6F6F6] dark:bg-zinc-800 col-start-1 col-end-2 row-start-1 row-end-7 p-5 w-32 fixed min-h-full">
            <Navigation className="flex flex-col items-center gap-10">
                <Navigation.Container className="flex flex-col gap-5">
                    {
                        DASHBOARD_LINKS.map(({id, title, url, icon}) => (
                            <Navigation.Item key={id} href={url}>{icon}</Navigation.Item>
                        ))
                    }
                </Navigation.Container>
                <Navigation.Container>
                    <ThemeToggle/>
                </Navigation.Container>
            </Navigation>
        </aside>
    )
}