import {Navigation, ThemeToggle} from "@/components";
import {DASHBOARD_LINKS} from "@/constants";

export const DashboardAside = () => {
    return (
        <aside className="bg-[#F6F6F6] dark:bg-zinc-800 min-h-full col-start-1 col-end-2 row-start-1 row-end-7 p-5">
            <Navigation className="flex flex-col content-center items-center justify-center gap-10">
                <Navigation.Container className="flex flex-col items-center gap-5">
                    {
                        DASHBOARD_LINKS.map(({id, title, url}) => (
                            <Navigation.Item key={id} href={url}>{title}</Navigation.Item>
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