import {Card} from "@/components";

export const HomeFeature = () => {
    return (
        <section className="container mx-auto py-5">
            <section className={"flex flex-col gap-10 text-center"}>
                <h2 className={"text-5xl font-bold text-center"}>New Features</h2>
                <p className={"text-2xl"}>Here are a few awesome features your going to love!</p>
            </section>
            <section className={"p-28 grid grid-cols-5 gap-5"}>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Authentication</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Subscriptions</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Payments</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Invoices</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Notifications</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Teams</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Permissions</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>Multi Tenant</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className={"text-2xl font-bold"}>User Profiles</h3>
                    </Card.Header>
                    <Card.Content>
                        lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam doloremque laudantium modi
                        pariatur reprehenderit!
                    </Card.Content>
                </Card>
            </section>
        </section>
    )
}