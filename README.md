# Gstore: A Data Store Specialized for Genomics

## Goal

The goal of the project is to design and implement a simple data store
specialized for genomics (huge two dimensional data). 

The basic idea is simple. The system stores data blobs in the two
different formats to support efficient lookup/scan in each dimension.
Suppose that we need three replicas for each data set (ignore erasure
coding), the system creates two replicas for row lookup and one
replica for column lookup. 

The combination of Bigtable and BigQuery (or HBase and Impara) can
achieve a similar goal. Compared with these systems, Gstore has the
following advantages:

- Minimize operational cost with its simplified architecture
- Store data efficiently by eliminating unnecessary data copy from a
  main key-value store to a columnar database

The target data are genome data, which are (mostly) immutable and
eventual consistency is sufficient. Durability is critical (i.e., no
data loss), but the availability target might not need to be very high
(e.g. 99.9% might be sufficient). These requirements also help us
simplify the design of the system.

## High Level Design

Let's consider the following basic example scenario:

1. Write data to (x_1, y), (x_2, y), ..., (x_n, y).
1. Read data from (x_1, y), (x_2, y), ..., (x_n, y).
1. Read data from (x, y_1), (x, y_2), ..., (x_n, y).

For genomics data, each row corresponds to genome position (i.e.,
locus), and each column corresponds to sample. (This is similar to what
[Hail](https://www.hail.is/docs/stable/overview.html#variant-dataset-vds)
has.)

The system conceptually has two instances of KV store (S1 and S2).
When writing v at (x, y), the system writes v at key K1 in S1 and key
K2 in S2 where K1 is `y/x` and K2 is `x/y`. Sine the KV store has data
locality based on key proximity (lexicographical order), scanning (x_i,
y) from S1 can be done efficiently. Scanning (x, y_i) from S2 can be
done efficiently for the same reason.

### Erasure Coding

Each genome is essentially an immutable blob encoded by
[Reed-Solomon](https://en.wikipedia.org/wiki/Reed%E2%80%93Solomon_error_correction).
Each genome is split into multiple "stripes". For each stripe, N data
chunks and M code chunks are created.

One interesting point here is that we need to write stripes to both S1
and S2, which have different data layouts. Given a stripe, its data
chunks are written to S1 as-is. On the other hand, the stripe needs
some transformation when it is written to S2.

Code chunks are stored in S1. Writes to S2 can be done later as long
as a sufficient number of data/code chunks are written to S1.

## Detailed Designs

### KV Store 

Simple master-slave. A design will be like https://github.com/kkaneda/blobstore/.

