import {DashboardAside, DashboardHeader} from "@/containers";

export default function Dashboard() {
    return (
        <div className={" grid grid-cols-12 grid-flow-row auto-rows-max gap-4"}>
            <DashboardHeader/>
            <DashboardAside/>
        </div>
    )
}