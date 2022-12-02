import {Navigation} from "@/components";
import {DASHBOARD_LINKS} from "@/constants";

export const DashboardAside = () => {
    return (
        <aside className="bg-amber-100">
            <Navigation>
                <Navigation.Container className={"flex flex-col gap-5"}>
                    {
                        DASHBOARD_LINKS.map(({id, title, url}) => (
                            <Navigation.Item key={id} href={url}>{title}</Navigation.Item>
                        ))
                    }
                </Navigation.Container>
            </Navigation>
        </aside>
    )
}