// mpi.go implements the core MPI-like primitives such as MPI_Init, MPI_Send, MPI_Recv,
// and MPI_Finalize. These functions allow distributed nodes to communicate by
// passing messages via AWS SQS and managing their rank and communication environment.
