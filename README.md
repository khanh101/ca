# ca
computational algebra

## pardic.go

implementation of $p$-adic integers

## uint1792.go

- my greatest appreciation to [apgoucher](https://cp4space.hatsya.com/2021/09/01/an-efficient-prime-for-number-theoretic-transforms/) to their prime $p = 2^{64} - 2^{32} + 1$ with $2$ being the $192$-th primitive root of unity in mod $p$
- `uint1792` with division $\lfloor a / b \rfloor$ for $1 < b < 2^{896}$ (can increase the bound close to $2^{1792}$)
- multiplication using FFT
- TODO : use mixed-radix Cooley-Tukey FFT
- TODO : use different primes then use CRT to construct output in larger base
