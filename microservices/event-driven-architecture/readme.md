## Event driven architecture
- Event occurs when a change happens in a system
- Highly distributed/loosely coupled architecture
- the services do not need to be aware of each other
- All listeners get notified of event and may take action

## Direct messaging
- services call other services directly

Advantages
- you don't need a seperate broker to communicate
- it has some advantages like easier implementation and error control
- if you call a service it can either give you an error immediately

Disadvantages
- increased coupling
- difficult to scale (thousands)

## Communication via message broker
Advantages
- this is more scalable
- increased decoupling

Disadvantages
- broker failure would be catastrophic

Communication patterns
- customer facing
 - you want to use request/response service
impact customers indirectly
 - use event driven model





