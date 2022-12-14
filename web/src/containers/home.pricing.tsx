import {Card} from "@/components";

export const HomePricing = () => {
    return (
        <section className="container mx-auto p-24 flex flex-col gap-10">
            <section className={"flex flex-col gap-10 text-center"}>
                <h2 className={"text-5xl font-bold text-center"}>Pricing plans</h2>
                <p className={"text-2xl"}>Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Aperiam doloremque laudantium modi pariatur reprehenderit!</p>
            </section>
            <section className="grid grid-cols-4 gap-10">
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">Small</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                            Aperiam</p>
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">Seed</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                            Aperiam</p>
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">Scale</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                            Aperiam</p>
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">Smart</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                            Aperiam</p>
                    </Card.Content>
                </Card>
            </section>
        </section>
    )
}