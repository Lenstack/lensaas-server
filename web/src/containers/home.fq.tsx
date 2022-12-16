import {Card} from "@/components";

export const HomeFq = () => {
    return (
        <section className="container mx-auto p-24 flex flex-col gap-10">
            <section className={"flex flex-col gap-10 text-center"}>
                <h2 className={"text-5xl font-bold text-center"}>Got questions?</h2>
                <p className={"text-2xl"}>Ask us anything concerning tech, pricing and licensing.</p>
            </section>
            <section className="grid grid-cols-2 gap-10">
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">What is Lorem Ipsum?</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                            Lorem Ipsum has been the industries standard dummy text ever since the 1500s,
                            when an unknown printer took a galley of type and scrambled it to make a type
                            specimen book. It has survived not only five centuries, but also the leap into
                            electronic typesetting, remaining essentially unchanged. It was popularised in
                            the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
                            and more recently with desktop publishing software like Aldus PageMaker
                            including versions of Lorem Ipsum.</p>
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">What is Lorem Ipsum?</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                            Lorem Ipsum has been the industries standard dummy text ever since the 1500s,
                            when an unknown printer took a galley of type and scrambled it to make a type
                            specimen book. It has survived not only five centuries, but also the leap into
                            electronic typesetting, remaining essentially unchanged. It was popularised in
                            the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
                            and more recently with desktop publishing software like Aldus PageMaker
                            including versions of Lorem Ipsum.</p>
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">What is Lorem Ipsum?</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                            Lorem Ipsum has been the industries standard dummy text ever since the 1500s,
                            when an unknown printer took a galley of type and scrambled it to make a type
                            specimen book. It has survived not only five centuries, but also the leap into
                            electronic typesetting, remaining essentially unchanged. It was popularised in
                            the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
                            and more recently with desktop publishing software like Aldus PageMaker
                            including versions of Lorem Ipsum.</p>
                    </Card.Content>
                </Card>
                <Card>
                    <Card.Header>
                        <h3 className="text-2xl font-bold">What is Lorem Ipsum?</h3>
                    </Card.Header>
                    <Card.Content>
                        <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                            Lorem Ipsum has been the industries standard dummy text ever since the 1500s,
                            when an unknown printer took a galley of type and scrambled it to make a type
                            specimen book. It has survived not only five centuries, but also the leap into
                            electronic typesetting, remaining essentially unchanged. It was popularised in
                            the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
                            and more recently with desktop publishing software like Aldus PageMaker
                            including versions of Lorem Ipsum.</p>
                    </Card.Content>
                </Card>
            </section>
        </section>
    )
}