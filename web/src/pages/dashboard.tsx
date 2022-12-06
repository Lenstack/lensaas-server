import {DashboardAside, DashboardHeader} from "@/containers";

export default function Dashboard() {
    return (
        <div className={"bg-gradient-to-b from-gray-100 to-gray-300 min-h-screen overflow-hidden"}>
            <DashboardHeader/>
            <DashboardAside/>
        </div>
    )
}