graph TD
    A[Main Context (ctx)] --> B[ErrGroup + Derived Context (group, ctx)]
    B --> C[Task 1]
    B --> D[Task 2 (Error)]
    B --> E[Task 3 (Cancelled)]
    D --> F[Error Propagation]
    F --> G[Group Wait]
    E --> G
    C --> G

    subgraph "Error Handling Flow"
        D --> F
    end

    subgraph "Group Wait"
        F --> G
        E --> G
        C --> G
    end

