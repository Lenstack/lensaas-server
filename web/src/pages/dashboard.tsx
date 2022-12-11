import {DashboardAside, DashboardHeader} from "@/containers";

export default function Dashboard() {
    return (
        <div className="min-h-screen overflow-hidden">
            <DashboardHeader/>
            <DashboardAside/>
        </div>
    )
}