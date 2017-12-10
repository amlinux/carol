// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

//------------------------------------------------------------------------------

import (
	"math"
	"testing"
)

//------------------------------------------------------------------------------

func TestSin(t *testing.T) {
	var x float32
	for _, tt := range sinTests {
		x = Sin(tt.in)
		if !IsNearlyEqual(x, tt.out, EpsilonFloat32) {
			t.Errorf("Relative error for Sin(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
		if !IsAlmostEqual(x, tt.out, 2) {
			t.Errorf("ULP error for Sin(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
	}
}

//------------------------------------------------------------------------------

func BenchmarkSin_math64(b *testing.B) {
	a := float64(0.5)
	for i := 0; i < b.N; i++ {
		_ = math.Sin(a)
	}
}

//------------------------------------------------------------------------------

func BenchmarkSin_math32(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = float32(math.Sin(float64(a)))
	}
}

//------------------------------------------------------------------------------

func BenchmarkSin_glam(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = Sin(a)
	}
}

//------------------------------------------------------------------------------

var sinTests = [...]struct {
	in  float32
	out float32
}{
	// Special cases

	{0, 0},
	{Pi / 6, 0.5},
	{Pi / 4, 0.7071067811865475244008443621048},
	{Pi / 3, 0.8660254037844386467637231707529},
	{Pi / 2, 1},

	// The following values have been generated with Gnu MPFR, 24 bits mantissa, roundTiesToEven

	{-3.14000010e+00, -1.59254798e-03},
	{-3.13000011e+00, -1.15922792e-02},
	{-3.12000012e+00, -2.15908512e-02},
	{-3.11000013e+00, -3.15872654e-02},
	{-3.10000014e+00, -4.15805206e-02},
	{-3.09000015e+00, -5.15696146e-02},
	{-3.08000016e+00, -6.15535565e-02},
	{-3.07000017e+00, -7.15313405e-02},
	{-3.06000018e+00, -8.15019682e-02},
	{-3.05000019e+00, -9.14644524e-02},
	{-3.04000020e+00, -1.01417787e-01},
	{-3.03000021e+00, -1.11360982e-01},
	{-3.02000022e+00, -1.21293038e-01},
	{-3.01000023e+00, -1.31212965e-01},
	{-3.00000024e+00, -1.41119778e-01},
	{-2.99000025e+00, -1.51012465e-01},
	{-2.98000026e+00, -1.60890058e-01},
	{-2.97000027e+00, -1.70751572e-01},
	{-2.96000028e+00, -1.80595994e-01},
	{-2.95000029e+00, -1.90422371e-01},
	{-2.94000030e+00, -2.00229689e-01},
	{-2.93000031e+00, -2.10016996e-01},
	{-2.92000031e+00, -2.19783306e-01},
	{-2.91000032e+00, -2.29527637e-01},
	{-2.90000033e+00, -2.39249006e-01},
	{-2.89000034e+00, -2.48946458e-01},
	{-2.88000035e+00, -2.58619010e-01},
	{-2.87000036e+00, -2.68265694e-01},
	{-2.86000037e+00, -2.77885556e-01},
	{-2.85000038e+00, -2.87477642e-01},
	{-2.84000039e+00, -2.97040969e-01},
	{-2.83000040e+00, -3.06574613e-01},
	{-2.82000041e+00, -3.16077590e-01},
	{-2.81000042e+00, -3.25548947e-01},
	{-2.80000043e+00, -3.34987760e-01},
	{-2.79000044e+00, -3.44393045e-01},
	{-2.78000045e+00, -3.53763938e-01},
	{-2.77000046e+00, -3.63099426e-01},
	{-2.76000047e+00, -3.72398615e-01},
	{-2.75000048e+00, -3.81660551e-01},
	{-2.74000049e+00, -3.90884340e-01},
	{-2.73000050e+00, -4.00069028e-01},
	{-2.72000051e+00, -4.09213722e-01},
	{-2.71000051e+00, -4.18317467e-01},
	{-2.70000052e+00, -4.27379400e-01},
	{-2.69000053e+00, -4.36398596e-01},
	{-2.68000054e+00, -4.45374161e-01},
	{-2.67000055e+00, -4.54305172e-01},
	{-2.66000056e+00, -4.63190764e-01},
	{-2.65000057e+00, -4.72030044e-01},
	{-2.64000058e+00, -4.80822116e-01},
	{-2.63000059e+00, -4.89566088e-01},
	{-2.62000060e+00, -4.98261124e-01},
	{-2.61000061e+00, -5.06906331e-01},
	{-2.60000062e+00, -5.15500844e-01},
	{-2.59000063e+00, -5.24043798e-01},
	{-2.58000064e+00, -5.32534361e-01},
	{-2.57000065e+00, -5.40971696e-01},
	{-2.56000066e+00, -5.49354911e-01},
	{-2.55000067e+00, -5.57683170e-01},
	{-2.54000068e+00, -5.65955698e-01},
	{-2.53000069e+00, -5.74171603e-01},
	{-2.52000070e+00, -5.82330108e-01},
	{-2.51000071e+00, -5.90430319e-01},
	{-2.50000072e+00, -5.98471582e-01},
	{-2.49000072e+00, -6.06452942e-01},
	{-2.48000073e+00, -6.14373684e-01},
	{-2.47000074e+00, -6.22232974e-01},
	{-2.46000075e+00, -6.30030036e-01},
	{-2.45000076e+00, -6.37764096e-01},
	{-2.44000077e+00, -6.45434380e-01},
	{-2.43000078e+00, -6.53040171e-01},
	{-2.42000079e+00, -6.60580635e-01},
	{-2.41000080e+00, -6.68054998e-01},
	{-2.40000081e+00, -6.75462604e-01},
	{-2.39000082e+00, -6.82802618e-01},
	{-2.38000083e+00, -6.90074384e-01},
	{-2.37000084e+00, -6.97277129e-01},
	{-2.36000085e+00, -7.04410136e-01},
	{-2.35000086e+00, -7.11472750e-01},
	{-2.34000087e+00, -7.18464196e-01},
	{-2.33000088e+00, -7.25383759e-01},
	{-2.32000089e+00, -7.32230842e-01},
	{-2.31000090e+00, -7.39004672e-01},
	{-2.30000091e+00, -7.45704591e-01},
	{-2.29000092e+00, -7.52329946e-01},
	{-2.28000093e+00, -7.58880079e-01},
	{-2.27000093e+00, -7.65354335e-01},
	{-2.26000094e+00, -7.71752059e-01},
	{-2.25000095e+00, -7.78072596e-01},
	{-2.24000096e+00, -7.84315348e-01},
	{-2.23000097e+00, -7.90479600e-01},
	{-2.22000098e+00, -7.96564877e-01},
	{-2.21000099e+00, -8.02570462e-01},
	{-2.20000100e+00, -8.08495820e-01},
	{-2.19000101e+00, -8.14340293e-01},
	{-2.18000102e+00, -8.20103347e-01},
	{-2.17000103e+00, -8.25784385e-01},
	{-2.16000104e+00, -8.31382871e-01},
	{-2.15000105e+00, -8.36898208e-01},
	{-2.14000106e+00, -8.42329860e-01},
	{-2.13000107e+00, -8.47677290e-01},
	{-2.12000108e+00, -8.52939904e-01},
	{-2.11000109e+00, -8.58117282e-01},
	{-2.10000110e+00, -8.63208830e-01},
	{-2.09000111e+00, -8.68214011e-01},
	{-2.08000112e+00, -8.73132408e-01},
	{-2.07000113e+00, -8.77963543e-01},
	{-2.06000113e+00, -8.82706821e-01},
	{-2.05000114e+00, -8.87361825e-01},
	{-2.04000115e+00, -8.91928136e-01},
	{-2.03000116e+00, -8.96405220e-01},
	{-2.02000117e+00, -9.00792658e-01},
	{-2.01000118e+00, -9.05090034e-01},
	{-2.00000119e+00, -9.09296930e-01},
	{-1.99000120e+00, -9.13412869e-01},
	{-1.98000121e+00, -9.17437494e-01},
	{-1.97000122e+00, -9.21370327e-01},
	{-1.96000123e+00, -9.25211072e-01},
	{-1.95000124e+00, -9.28959250e-01},
	{-1.94000125e+00, -9.32614565e-01},
	{-1.93000126e+00, -9.36176598e-01},
	{-1.92000127e+00, -9.39645052e-01},
	{-1.91000128e+00, -9.43019509e-01},
	{-1.90000129e+00, -9.46299672e-01},
	{-1.89000130e+00, -9.49485183e-01},
	{-1.88000131e+00, -9.52575803e-01},
	{-1.87000132e+00, -9.55571115e-01},
	{-1.86000133e+00, -9.58470881e-01},
	{-1.85000134e+00, -9.61274862e-01},
	{-1.84000134e+00, -9.63982642e-01},
	{-1.83000135e+00, -9.66594040e-01},
	{-1.82000136e+00, -9.69108820e-01},
	{-1.81000137e+00, -9.71526623e-01},
	{-1.80000138e+00, -9.73847330e-01},
	{-1.79000139e+00, -9.76070642e-01},
	{-1.78000140e+00, -9.78196323e-01},
	{-1.77000141e+00, -9.80224192e-01},
	{-1.76000142e+00, -9.82154071e-01},
	{-1.75000143e+00, -9.83985662e-01},
	{-1.74000144e+00, -9.85718966e-01},
	{-1.73000145e+00, -9.87353623e-01},
	{-1.72000146e+00, -9.88889575e-01},
	{-1.71000147e+00, -9.90326583e-01},
	{-1.70000148e+00, -9.91664648e-01},
	{-1.69000149e+00, -9.92903471e-01},
	{-1.68000150e+00, -9.94043052e-01},
	{-1.67000151e+00, -9.95083213e-01},
	{-1.66000152e+00, -9.96023834e-01},
	{-1.65000153e+00, -9.96864915e-01},
	{-1.64000154e+00, -9.97606277e-01},
	{-1.63000154e+00, -9.98247862e-01},
	{-1.62000155e+00, -9.98789668e-01},
	{-1.61000156e+00, -9.99231577e-01},
	{-1.60000157e+00, -9.99573529e-01},
	{-1.59000158e+00, -9.99815583e-01},
	{-1.58000159e+00, -9.99957621e-01},
	{-1.57000160e+00, -9.99999702e-01},
	{-1.56000161e+00, -9.99941766e-01},
	{-1.55000162e+00, -9.99783814e-01},
	{-1.54000163e+00, -9.99525905e-01},
	{-1.53000164e+00, -9.99168038e-01},
	{-1.52000165e+00, -9.98710215e-01},
	{-1.51000166e+00, -9.98152554e-01},
	{-1.50000167e+00, -9.97495115e-01},
	{-1.49000168e+00, -9.96737897e-01},
	{-1.48000169e+00, -9.95881021e-01},
	{-1.47000170e+00, -9.94924545e-01},
	{-1.46000171e+00, -9.93868530e-01},
	{-1.45000172e+00, -9.92713213e-01},
	{-1.44000173e+00, -9.91458595e-01},
	{-1.43000174e+00, -9.90104795e-01},
	{-1.42000175e+00, -9.88652050e-01},
	{-1.41000175e+00, -9.87100363e-01},
	{-1.40000176e+00, -9.85450029e-01},
	{-1.39000177e+00, -9.83701110e-01},
	{-1.38000178e+00, -9.81853843e-01},
	{-1.37000179e+00, -9.79908407e-01},
	{-1.36000180e+00, -9.77864981e-01},
	{-1.35000181e+00, -9.75723743e-01},
	{-1.34000182e+00, -9.73484933e-01},
	{-1.33000183e+00, -9.71148789e-01},
	{-1.32000184e+00, -9.68715549e-01},
	{-1.31000185e+00, -9.66185451e-01},
	{-1.30000186e+00, -9.63558674e-01},
	{-1.29000187e+00, -9.60835576e-01},
	{-1.28000188e+00, -9.58016396e-01},
	{-1.27000189e+00, -9.55101430e-01},
	{-1.26000190e+00, -9.52090919e-01},
	{-1.25000191e+00, -9.48985219e-01},
	{-1.24000192e+00, -9.45784628e-01},
	{-1.23000193e+00, -9.42489445e-01},
	{-1.22000194e+00, -9.39100027e-01},
	{-1.21000195e+00, -9.35616672e-01},
	{-1.20000196e+00, -9.32039797e-01},
	{-1.19000196e+00, -9.28369701e-01},
	{-1.18000197e+00, -9.24606740e-01},
	{-1.17000198e+00, -9.20751393e-01},
	{-1.16000199e+00, -9.16803896e-01},
	{-1.15000200e+00, -9.12764788e-01},
	{-1.14000201e+00, -9.08634365e-01},
	{-1.13000202e+00, -9.04413044e-01},
	{-1.12000203e+00, -9.00101304e-01},
	{-1.11000204e+00, -8.95699620e-01},
	{-1.10000205e+00, -8.91208291e-01},
	{-1.09000206e+00, -8.86627853e-01},
	{-1.08000207e+00, -8.81958783e-01},
	{-1.07000208e+00, -8.77201498e-01},
	{-1.06000209e+00, -8.72356474e-01},
	{-1.05000210e+00, -8.67424250e-01},
	{-1.04000211e+00, -8.62405300e-01},
	{-1.03000212e+00, -8.57300103e-01},
	{-1.02000213e+00, -8.52109134e-01},
	{-1.01000214e+00, -8.46832991e-01},
	{-1.00000215e+00, -8.41472149e-01},
	{-9.90002155e-01, -8.36027145e-01},
	{-9.80002165e-01, -8.30498576e-01},
	{-9.70002174e-01, -8.24886918e-01},
	{-9.60002184e-01, -8.19192827e-01},
	{-9.50002193e-01, -8.13416779e-01},
	{-9.40002203e-01, -8.07559371e-01},
	{-9.30002213e-01, -8.01621258e-01},
	{-9.20002222e-01, -7.95602977e-01},
	{-9.10002232e-01, -7.89505124e-01},
	{-9.00002241e-01, -7.83328295e-01},
	{-8.90002251e-01, -7.77073145e-01},
	{-8.80002260e-01, -7.70740330e-01},
	{-8.70002270e-01, -7.64330387e-01},
	{-8.60002279e-01, -7.57844031e-01},
	{-8.50002289e-01, -7.51281917e-01},
	{-8.40002298e-01, -7.44644642e-01},
	{-8.30002308e-01, -7.37932920e-01},
	{-8.20002317e-01, -7.31147408e-01},
	{-8.10002327e-01, -7.24288762e-01},
	{-8.00002337e-01, -7.17357695e-01},
	{-7.90002346e-01, -7.10354924e-01},
	{-7.80002356e-01, -7.03281105e-01},
	{-7.70002365e-01, -6.96136951e-01},
	{-7.60002375e-01, -6.88923180e-01},
	{-7.50002384e-01, -6.81640506e-01},
	{-7.40002394e-01, -6.74289703e-01},
	{-7.30002403e-01, -6.66871428e-01},
	{-7.20002413e-01, -6.59386516e-01},
	{-7.10002422e-01, -6.51835620e-01},
	{-7.00002432e-01, -6.44219518e-01},
	{-6.90002441e-01, -6.36539042e-01},
	{-6.80002451e-01, -6.28794909e-01},
	{-6.70002460e-01, -6.20987892e-01},
	{-6.60002470e-01, -6.13118827e-01},
	{-6.50002480e-01, -6.05188370e-01},
	{-6.40002489e-01, -5.97197413e-01},
	{-6.30002499e-01, -5.89146793e-01},
	{-6.20002508e-01, -5.81037223e-01},
	{-6.10002518e-01, -5.72869539e-01},
	{-6.00002527e-01, -5.64644575e-01},
	{-5.90002537e-01, -5.56363106e-01},
	{-5.80002546e-01, -5.48026085e-01},
	{-5.70002556e-01, -5.39634228e-01},
	{-5.60002565e-01, -5.31188369e-01},
	{-5.50002575e-01, -5.22689402e-01},
	{-5.40002584e-01, -5.14138222e-01},
	{-5.30002594e-01, -5.05535603e-01},
	{-5.20002604e-01, -4.96882409e-01},
	{-5.10002613e-01, -4.88179535e-01},
	{-5.00002623e-01, -4.79427844e-01},
	{-4.90002632e-01, -4.70628202e-01},
	{-4.80002642e-01, -4.61781532e-01},
	{-4.70002651e-01, -4.52888638e-01},
	{-4.60002661e-01, -4.43950504e-01},
	{-4.50002670e-01, -4.34967935e-01},
	{-4.40002680e-01, -4.25941885e-01},
	{-4.30002689e-01, -4.16873246e-01},
	{-4.20002699e-01, -4.07762915e-01},
	{-4.10002708e-01, -3.98611814e-01},
	{-4.00002718e-01, -3.89420837e-01},
	{-3.90002728e-01, -3.80190939e-01},
	{-3.80002737e-01, -3.70923012e-01},
	{-3.70002747e-01, -3.61617982e-01},
	{-3.60002756e-01, -3.52276802e-01},
	{-3.50002766e-01, -3.42900395e-01},
	{-3.40002775e-01, -3.33489716e-01},
	{-3.30002785e-01, -3.24045658e-01},
	{-3.20002794e-01, -3.14569205e-01},
	{-3.10002804e-01, -3.05061311e-01},
	{-3.00002813e-01, -2.95522898e-01},
	{-2.90002823e-01, -2.85954922e-01},
	{-2.80002832e-01, -2.76358366e-01},
	{-2.70002842e-01, -2.66734183e-01},
	{-2.60002851e-01, -2.57083297e-01},
	{-2.50002861e-01, -2.47406736e-01},
	{-2.40002856e-01, -2.37705395e-01},
	{-2.30002850e-01, -2.27980301e-01},
	{-2.20002845e-01, -2.18232393e-01},
	{-2.10002840e-01, -2.08462670e-01},
	{-2.00002834e-01, -1.98672116e-01},
	{-1.90002829e-01, -1.88861668e-01},
	{-1.80002823e-01, -1.79032356e-01},
	{-1.70002818e-01, -1.69185132e-01},
	{-1.60002813e-01, -1.59320980e-01},
	{-1.50002807e-01, -1.49440914e-01},
	{-1.40002802e-01, -1.39545888e-01},
	{-1.30002797e-01, -1.29636914e-01},
	{-1.20002799e-01, -1.19714983e-01},
	{-1.10002801e-01, -1.09781086e-01},
	{-1.00002803e-01, -9.98362079e-02},
	{-9.00028050e-02, -8.98813456e-02},
	{-8.00028071e-02, -7.99174905e-02},
	{-7.00028092e-02, -6.99456483e-02},
	{-6.00028113e-02, -5.99668138e-02},
	{-5.00028133e-02, -4.99819778e-02},
	{-4.00028154e-02, -3.99921462e-02},
	{-3.00028156e-02, -2.99983155e-02},
	{-2.00028159e-02, -2.00014822e-02},
	{-1.00028161e-02, -1.00026494e-02},
	{-2.81631947e-06, -2.81631947e-06},
	{+9.99718346e-03, +9.99701675e-03},
	{+1.99971832e-02, +1.99958496e-02},
	{+2.99971830e-02, +2.99926847e-02},
	{+3.99971828e-02, +3.99865210e-02},
	{+4.99971807e-02, +4.99763526e-02},
	{+5.99971786e-02, +5.99611886e-02},
	{+6.99971765e-02, +6.99400306e-02},
	{+7.99971744e-02, +7.99118802e-02},
	{+8.99971724e-02, +8.98757353e-02},
	{+9.99971703e-02, +9.98305976e-02},
	{+1.09997168e-01, +1.09775484e-01},
	{+1.19997166e-01, +1.19709395e-01},
	{+1.29997164e-01, +1.29631326e-01},
	{+1.39997169e-01, +1.39540315e-01},
	{+1.49997175e-01, +1.49435341e-01},
	{+1.59997180e-01, +1.59315422e-01},
	{+1.69997185e-01, +1.69179574e-01},
	{+1.79997191e-01, +1.79026812e-01},
	{+1.89997196e-01, +1.88856140e-01},
	{+1.99997202e-01, +1.98666587e-01},
	{+2.09997207e-01, +2.08457172e-01},
	{+2.19997212e-01, +2.18226910e-01},
	{+2.29997218e-01, +2.27974817e-01},
	{+2.39997223e-01, +2.37699926e-01},
	{+2.49997228e-01, +2.47401267e-01},
	{+2.59997219e-01, +2.57077873e-01},
	{+2.69997209e-01, +2.66728759e-01},
	{+2.79997200e-01, +2.76352972e-01},
	{+2.89997190e-01, +2.85949528e-01},
	{+2.99997181e-01, +2.95517504e-01},
	{+3.09997171e-01, +3.05055946e-01},
	{+3.19997162e-01, +3.14563870e-01},
	{+3.29997152e-01, +3.24040323e-01},
	{+3.39997143e-01, +3.33484411e-01},
	{+3.49997133e-01, +3.42895120e-01},
	{+3.59997123e-01, +3.52271527e-01},
	{+3.69997114e-01, +3.61612737e-01},
	{+3.79997104e-01, +3.70917767e-01},
	{+3.89997095e-01, +3.80185723e-01},
	{+3.99997085e-01, +3.89415652e-01},
	{+4.09997076e-01, +3.98606658e-01},
	{+4.19997066e-01, +4.07757789e-01},
	{+4.29997057e-01, +4.16868120e-01},
	{+4.39997047e-01, +4.25936788e-01},
	{+4.49997038e-01, +4.34962869e-01},
	{+4.59997028e-01, +4.43945438e-01},
	{+4.69997019e-01, +4.52883631e-01},
	{+4.79997009e-01, +4.61776525e-01},
	{+4.89997000e-01, +4.70623255e-01},
	{+4.99996990e-01, +4.79422897e-01},
	{+5.09997010e-01, +4.88174647e-01},
	{+5.19997001e-01, +4.96877521e-01},
	{+5.29996991e-01, +5.05530775e-01},
	{+5.39996982e-01, +5.14133394e-01},
	{+5.49996972e-01, +5.22684634e-01},
	{+5.59996963e-01, +5.31183600e-01},
	{+5.69996953e-01, +5.39629459e-01},
	{+5.79996943e-01, +5.48021376e-01},
	{+5.89996934e-01, +5.56358457e-01},
	{+5.99996924e-01, +5.64639926e-01},
	{+6.09996915e-01, +5.72864950e-01},
	{+6.19996905e-01, +5.81032634e-01},
	{+6.29996896e-01, +5.89142263e-01},
	{+6.39996886e-01, +5.97192943e-01},
	{+6.49996877e-01, +6.05183899e-01},
	{+6.59996867e-01, +6.13114357e-01},
	{+6.69996858e-01, +6.20983541e-01},
	{+6.79996848e-01, +6.28790557e-01},
	{+6.89996839e-01, +6.36534750e-01},
	{+6.99996829e-01, +6.44215286e-01},
	{+7.09996819e-01, +6.51831388e-01},
	{+7.19996810e-01, +6.59382284e-01},
	{+7.29996800e-01, +6.66867256e-01},
	{+7.39996791e-01, +6.74285531e-01},
	{+7.49996781e-01, +6.81636393e-01},
	{+7.59996772e-01, +6.88919127e-01},
	{+7.69996762e-01, +6.96132898e-01},
	{+7.79996753e-01, +7.03277111e-01},
	{+7.89996743e-01, +7.10350990e-01},
	{+7.99996734e-01, +7.17353821e-01},
	{+8.09996724e-01, +7.24284887e-01},
	{+8.19996715e-01, +7.31143594e-01},
	{+8.29996705e-01, +7.37929165e-01},
	{+8.39996696e-01, +7.44640887e-01},
	{+8.49996686e-01, +7.51278222e-01},
	{+8.59996676e-01, +7.57840395e-01},
	{+8.69996667e-01, +7.64326811e-01},
	{+8.79996657e-01, +7.70736754e-01},
	{+8.89996648e-01, +7.77069628e-01},
	{+8.99996638e-01, +7.83324838e-01},
	{+9.09996629e-01, +7.89501667e-01},
	{+9.19996619e-01, +7.95599580e-01},
	{+9.29996610e-01, +8.01617920e-01},
	{+9.39996600e-01, +8.07556093e-01},
	{+9.49996591e-01, +8.13413501e-01},
	{+9.59996581e-01, +8.19189608e-01},
	{+9.69996572e-01, +8.24883759e-01},
	{+9.79996562e-01, +8.30495477e-01},
	{+9.89996552e-01, +8.36024106e-01},
	{+9.99996543e-01, +8.41469109e-01},
	{+1.00999653e+00, +8.46830010e-01},
	{+1.01999652e+00, +8.52106214e-01},
	{+1.02999651e+00, +8.57297182e-01},
	{+1.03999650e+00, +8.62402439e-01},
	{+1.04999650e+00, +8.67421508e-01},
	{+1.05999649e+00, +8.72353792e-01},
	{+1.06999648e+00, +8.77198815e-01},
	{+1.07999647e+00, +8.81956160e-01},
	{+1.08999646e+00, +8.86625290e-01},
	{+1.09999645e+00, +8.91205728e-01},
	{+1.10999644e+00, +8.95697117e-01},
	{+1.11999643e+00, +9.00098860e-01},
	{+1.12999642e+00, +9.04410660e-01},
	{+1.13999641e+00, +9.08631980e-01},
	{+1.14999640e+00, +9.12762463e-01},
	{+1.15999639e+00, +9.16801691e-01},
	{+1.16999638e+00, +9.20749187e-01},
	{+1.17999637e+00, +9.24604654e-01},
	{+1.18999636e+00, +9.28367615e-01},
	{+1.19999635e+00, +9.32037771e-01},
	{+1.20999634e+00, +9.35614705e-01},
	{+1.21999633e+00, +9.39098120e-01},
	{+1.22999632e+00, +9.42487597e-01},
	{+1.23999631e+00, +9.45782781e-01},
	{+1.24999630e+00, +9.48983431e-01},
	{+1.25999629e+00, +9.52089190e-01},
	{+1.26999629e+00, +9.55099761e-01},
	{+1.27999628e+00, +9.58014786e-01},
	{+1.28999627e+00, +9.60834026e-01},
	{+1.29999626e+00, +9.63557184e-01},
	{+1.30999625e+00, +9.66183960e-01},
	{+1.31999624e+00, +9.68714178e-01},
	{+1.32999623e+00, +9.71147478e-01},
	{+1.33999622e+00, +9.73483682e-01},
	{+1.34999621e+00, +9.75722551e-01},
	{+1.35999620e+00, +9.77863789e-01},
	{+1.36999619e+00, +9.79907274e-01},
	{+1.37999618e+00, +9.81852829e-01},
	{+1.38999617e+00, +9.83700097e-01},
	{+1.39999616e+00, +9.85449076e-01},
	{+1.40999615e+00, +9.87099469e-01},
	{+1.41999614e+00, +9.88651156e-01},
	{+1.42999613e+00, +9.90104020e-01},
	{+1.43999612e+00, +9.91457820e-01},
	{+1.44999611e+00, +9.92712498e-01},
	{+1.45999610e+00, +9.93867934e-01},
	{+1.46999609e+00, +9.94923949e-01},
	{+1.47999609e+00, +9.95880485e-01},
	{+1.48999608e+00, +9.96737421e-01},
	{+1.49999607e+00, +9.97494698e-01},
	{+1.50999606e+00, +9.98152256e-01},
	{+1.51999605e+00, +9.98709917e-01},
	{+1.52999604e+00, +9.99167800e-01},
	{+1.53999603e+00, +9.99525726e-01},
	{+1.54999602e+00, +9.99783695e-01},
	{+1.55999601e+00, +9.99941707e-01},
	{+1.56999600e+00, +9.99999702e-01},
	{+1.57999599e+00, +9.99957681e-01},
	{+1.58999598e+00, +9.99815702e-01},
	{+1.59999597e+00, +9.99573708e-01},
	{+1.60999596e+00, +9.99231815e-01},
	{+1.61999595e+00, +9.98789966e-01},
	{+1.62999594e+00, +9.98248219e-01},
	{+1.63999593e+00, +9.97606635e-01},
	{+1.64999592e+00, +9.96865332e-01},
	{+1.65999591e+00, +9.96024370e-01},
	{+1.66999590e+00, +9.95083749e-01},
	{+1.67999589e+00, +9.94043648e-01},
	{+1.68999588e+00, +9.92904127e-01},
	{+1.69999588e+00, +9.91665363e-01},
	{+1.70999587e+00, +9.90327358e-01},
	{+1.71999586e+00, +9.88890409e-01},
	{+1.72999585e+00, +9.87354517e-01},
	{+1.73999584e+00, +9.85719860e-01},
	{+1.74999583e+00, +9.83986676e-01},
	{+1.75999582e+00, +9.82155085e-01},
	{+1.76999581e+00, +9.80225325e-01},
	{+1.77999580e+00, +9.78197455e-01},
	{+1.78999579e+00, +9.76071835e-01},
	{+1.79999578e+00, +9.73848581e-01},
	{+1.80999577e+00, +9.71527934e-01},
	{+1.81999576e+00, +9.69110191e-01},
	{+1.82999575e+00, +9.66595471e-01},
	{+1.83999574e+00, +9.63984132e-01},
	{+1.84999573e+00, +9.61276352e-01},
	{+1.85999572e+00, +9.58472490e-01},
	{+1.86999571e+00, +9.55572784e-01},
	{+1.87999570e+00, +9.52577531e-01},
	{+1.88999569e+00, +9.49486971e-01},
	{+1.89999568e+00, +9.46301460e-01},
	{+1.90999568e+00, +9.43021357e-01},
	{+1.91999567e+00, +9.39646959e-01},
	{+1.92999566e+00, +9.36178565e-01},
	{+1.93999565e+00, +9.32616591e-01},
	{+1.94999564e+00, +9.28961337e-01},
	{+1.95999563e+00, +9.25213158e-01},
	{+1.96999562e+00, +9.21372533e-01},
	{+1.97999561e+00, +9.17439699e-01},
	{+1.98999560e+00, +9.13415134e-01},
	{+1.99999559e+00, +9.09299254e-01},
	{+2.00999570e+00, +9.05092418e-01},
	{+2.01999569e+00, +9.00795043e-01},
	{+2.02999568e+00, +8.96407664e-01},
	{+2.03999567e+00, +8.91930580e-01},
	{+2.04999566e+00, +8.87364388e-01},
	{+2.05999565e+00, +8.82709384e-01},
	{+2.06999564e+00, +8.77966166e-01},
	{+2.07999563e+00, +8.73135090e-01},
	{+2.08999562e+00, +8.68216753e-01},
	{+2.09999561e+00, +8.63211572e-01},
	{+2.10999560e+00, +8.58120084e-01},
	{+2.11999559e+00, +8.52942765e-01},
	{+2.12999558e+00, +8.47680211e-01},
	{+2.13999557e+00, +8.42332840e-01},
	{+2.14999557e+00, +8.36901248e-01},
	{+2.15999556e+00, +8.31385911e-01},
	{+2.16999555e+00, +8.25787485e-01},
	{+2.17999554e+00, +8.20106506e-01},
	{+2.18999553e+00, +8.14343512e-01},
	{+2.19999552e+00, +8.08499038e-01},
	{+2.20999551e+00, +8.02573740e-01},
	{+2.21999550e+00, +7.96568215e-01},
	{+2.22999549e+00, +7.90482998e-01},
	{+2.23999548e+00, +7.84318745e-01},
	{+2.24999547e+00, +7.78076053e-01},
	{+2.25999546e+00, +7.71755576e-01},
	{+2.26999545e+00, +7.65357852e-01},
	{+2.27999544e+00, +7.58883655e-01},
	{+2.28999543e+00, +7.52333581e-01},
	{+2.29999542e+00, +7.45708287e-01},
	{+2.30999541e+00, +7.39008367e-01},
	{+2.31999540e+00, +7.32234597e-01},
	{+2.32999539e+00, +7.25387573e-01},
	{+2.33999538e+00, +7.18468010e-01},
	{+2.34999537e+00, +7.11476624e-01},
	{+2.35999537e+00, +7.04414070e-01},
	{+2.36999536e+00, +6.97281063e-01},
	{+2.37999535e+00, +6.90078378e-01},
	{+2.38999534e+00, +6.82806611e-01},
	{+2.39999533e+00, +6.75466597e-01},
	{+2.40999532e+00, +6.68059051e-01},
	{+2.41999531e+00, +6.60584748e-01},
	{+2.42999530e+00, +6.53044283e-01},
	{+2.43999529e+00, +6.45438612e-01},
	{+2.44999528e+00, +6.37768328e-01},
	{+2.45999527e+00, +6.30034328e-01},
	{+2.46999526e+00, +6.22237265e-01},
	{+2.47999525e+00, +6.14377975e-01},
	{+2.48999524e+00, +6.06457293e-01},
	{+2.49999523e+00, +5.98475993e-01},
	{+2.50999522e+00, +5.90434790e-01},
	{+2.51999521e+00, +5.82334518e-01},
	{+2.52999520e+00, +5.74176073e-01},
	{+2.53999519e+00, +5.65960169e-01},
	{+2.54999518e+00, +5.57687700e-01},
	{+2.55999517e+00, +5.49359441e-01},
	{+2.56999516e+00, +5.40976286e-01},
	{+2.57999516e+00, +5.32539010e-01},
	{+2.58999515e+00, +5.24048448e-01},
	{+2.59999514e+00, +5.15505552e-01},
	{+2.60999513e+00, +5.06911039e-01},
	{+2.61999512e+00, +4.98265862e-01},
	{+2.62999511e+00, +4.89570886e-01},
	{+2.63999510e+00, +4.80826914e-01},
	{+2.64999509e+00, +4.72034872e-01},
	{+2.65999508e+00, +4.63195622e-01},
	{+2.66999507e+00, +4.54310060e-01},
	{+2.67999506e+00, +4.45379078e-01},
	{+2.68999505e+00, +4.36403543e-01},
	{+2.69999504e+00, +4.27384377e-01},
	{+2.70999503e+00, +4.18322444e-01},
	{+2.71999502e+00, +4.09218699e-01},
	{+2.72999501e+00, +4.00074035e-01},
	{+2.73999500e+00, +3.90889376e-01},
	{+2.74999499e+00, +3.81665617e-01},
	{+2.75999498e+00, +3.72403681e-01},
	{+2.76999497e+00, +3.63104522e-01},
	{+2.77999496e+00, +3.53769064e-01},
	{+2.78999496e+00, +3.44398201e-01},
	{+2.79999495e+00, +3.34992915e-01},
	{+2.80999494e+00, +3.25554132e-01},
	{+2.81999493e+00, +3.16082776e-01},
	{+2.82999492e+00, +3.06579828e-01},
	{+2.83999491e+00, +2.97046214e-01},
	{+2.84999490e+00, +2.87482888e-01},
	{+2.85999489e+00, +2.77890831e-01},
	{+2.86999488e+00, +2.68270999e-01},
	{+2.87999487e+00, +2.58624315e-01},
	{+2.88999486e+00, +2.48951763e-01},
	{+2.89999485e+00, +2.39254326e-01},
	{+2.90999484e+00, +2.29532972e-01},
	{+2.91999483e+00, +2.19788656e-01},
	{+2.92999482e+00, +2.10022360e-01},
	{+2.93999481e+00, +2.00235069e-01},
	{+2.94999480e+00, +1.90427750e-01},
	{+2.95999479e+00, +1.80601388e-01},
	{+2.96999478e+00, +1.70756966e-01},
	{+2.97999477e+00, +1.60895467e-01},
	{+2.98999476e+00, +1.51017889e-01},
	{+2.99999475e+00, +1.41125202e-01},
	{+3.00999475e+00, +1.31218404e-01},
	{+3.01999474e+00, +1.21298477e-01},
	{+3.02999473e+00, +1.11366428e-01},
	{+3.03999472e+00, +1.01423241e-01},
	{+3.04999471e+00, +9.14699137e-02},
	{+3.05999470e+00, +8.15074369e-02},
	{+3.06999469e+00, +7.15368092e-02},
	{+3.07999468e+00, +6.15590289e-02},
	{+3.08999467e+00, +5.15750907e-02},
	{+3.09999466e+00, +4.15859967e-02},
	{+3.10999465e+00, +3.15927453e-02},
	{+3.11999464e+00, +2.15963349e-02},
	{+3.12999463e+00, +1.15977628e-02},
	{+3.13999462e+00, +1.59803161e-03},
}

//------------------------------------------------------------------------------
