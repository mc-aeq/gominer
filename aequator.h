#ifndef __AEQUATOR_H__
#define __AEQUATOR_H__

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

void	aequator_hash_nonce(uint32_t grid, uint32_t block, uint32_t threads,
	    uint32_t startNonce, uint32_t *resNonce, uint32_t targetHigh);
void	aequator_cpu_setBlock_52(const uint32_t *input);

#ifdef __cplusplus
}
#endif /* __cplusplus */

#endif /* __AEQUATOR_H__ */
