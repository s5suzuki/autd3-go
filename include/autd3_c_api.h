﻿// File: autd3_c_api.h
// Project: base
// Created Date: 16/05/2022
// Author: Shun Suzuki
// -----
// Last Modified: 10/06/2022
// Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
// -----
// Copyright (c) 2022 Shun Suzuki. All rights reserved.
//

#pragma once

#include "./header.hpp"

#ifdef __cplusplus
extern "C" {
#endif
EXPORT_AUTD int32_t AUTDGetLastError(OUT char* error);
EXPORT_AUTD void AUTDCreateController(OUT void** out);
EXPORT_AUTD bool AUTDOpenController(IN void* handle, IN void* link);
EXPORT_AUTD int32_t AUTDAddDevice(IN void* handle, IN double x, IN double y, IN double z, IN double rz1, IN double ry, IN double rz2);
EXPORT_AUTD int32_t AUTDAddDeviceQuaternion(IN void* handle, IN double x, IN double y, IN double z, IN double qw, IN double qx, IN double qy,
                                            IN double qz);
EXPORT_AUTD int32_t AUTDClose(IN void* handle);
EXPORT_AUTD int32_t AUTDClear(IN void* handle);
EXPORT_AUTD int32_t AUTDSynchronize(IN void* handle);
EXPORT_AUTD void AUTDFreeController(IN const void* handle);
EXPORT_AUTD bool AUTDIsOpen(IN const void* handle);
EXPORT_AUTD bool AUTDGetForceFan(IN const void* handle);
EXPORT_AUTD bool AUTDGetReadsFPGAInfo(IN const void* handle);
EXPORT_AUTD bool AUTDGetCheckAck(IN const void* handle);
EXPORT_AUTD void AUTDSetReadsFPGAInfo(IN void* handle, IN bool reads_fpga_info);
EXPORT_AUTD void AUTDSetCheckAck(IN void* handle, IN bool check_ack);
EXPORT_AUTD void AUTDSetForceFan(IN void* handle, IN bool force);
EXPORT_AUTD double AUTDGetSoundSpeed(IN const void* handle);
EXPORT_AUTD void AUTDSetSoundSpeed(IN void* handle, IN double sound_speed);
EXPORT_AUTD double AUTDGetTransFrequency(IN const void* handle, IN int32_t device_idx, IN int32_t local_trans_idx);
EXPORT_AUTD void AUTDSetTransFrequency(IN void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, IN double frequency);
EXPORT_AUTD uint16_t AUTDGetTransCycle(IN const void* handle, IN int32_t device_idx, IN int32_t local_trans_idx);
EXPORT_AUTD void AUTDSetTransCycle(IN void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, IN uint16_t cycle);
EXPORT_AUTD double AUTDGetWavelength(IN const void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, IN double sound_speed);
EXPORT_AUTD double AUTDGetAttenuation(IN const void* handle);
EXPORT_AUTD void AUTDSetAttenuation(IN void* handle, IN double attenuation);
EXPORT_AUTD bool AUTDGetFPGAInfo(IN void* handle, IN uint8_t* out);
EXPORT_AUTD int32_t AUTDUpdateFlags(IN void* handle);
EXPORT_AUTD int32_t AUTDNumDevices(IN const void* handle);
EXPORT_AUTD void AUTDTransPosition(IN const void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, OUT double* x, OUT double* y,
                                   OUT double* z);
EXPORT_AUTD void AUTDTransXDirection(IN const void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, OUT double* x, OUT double* y,
                                     OUT double* z);
EXPORT_AUTD void AUTDTransYDirection(IN const void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, OUT double* x, OUT double* y,
                                     OUT double* z);
EXPORT_AUTD void AUTDTransZDirection(IN const void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, OUT double* x, OUT double* y,
                                     OUT double* z);
EXPORT_AUTD int32_t AUTDGetFirmwareInfoListPointer(IN void* handle, OUT void** out);
EXPORT_AUTD void AUTDGetFirmwareInfo(IN const void* p_firm_info_list, IN int32_t index, OUT char* info);
EXPORT_AUTD void AUTDFreeFirmwareInfoListPointer(IN const void* p_firm_info_list);
EXPORT_AUTD void AUTDGainNull(OUT void** gain);
EXPORT_AUTD void AUTDGainGrouped(OUT void** gain, IN const void* handle);
EXPORT_AUTD void AUTDGainGroupedAdd(IN void* grouped_gain, IN int32_t device_id, IN void* gain);
EXPORT_AUTD void AUTDGainFocus(OUT void** gain, IN double x, IN double y, IN double z, IN double amp);
EXPORT_AUTD void AUTDGainBesselBeam(OUT void** gain, IN double x, IN double y, IN double z, IN double n_x, IN double n_y, IN double n_z,
                                    IN double theta_z, IN double amp);
EXPORT_AUTD void AUTDGainPlaneWave(OUT void** gain, IN double n_x, IN double n_y, IN double n_z, IN double amp);
EXPORT_AUTD void AUTDGainCustom(OUT void** gain, IN const double* amp, IN const double* phase, IN uint64_t size);
EXPORT_AUTD void AUTDDeleteGain(IN const void* gain);
EXPORT_AUTD void AUTDModulationStatic(OUT void** mod, IN double amp);
EXPORT_AUTD void AUTDModulationSine(OUT void** mod, IN int32_t freq, IN double amp, IN double offset);
EXPORT_AUTD void AUTDModulationSineSquared(OUT void** mod, IN int32_t freq, IN double amp, IN double offset);
EXPORT_AUTD void AUTDModulationSineLegacy(OUT void** mod, IN double freq, IN double amp, IN double offset);
EXPORT_AUTD void AUTDModulationSquare(OUT void** mod, IN int32_t freq, IN double low, IN double high, IN double duty);
EXPORT_AUTD void AUTDModulationCustom(OUT void** mod, IN const uint8_t* buffer, IN uint64_t size, IN uint32_t freq_div);
EXPORT_AUTD uint32_t AUTDModulationSamplingFrequencyDivision(IN const void* mod);
EXPORT_AUTD void AUTDModulationSetSamplingFrequencyDivision(IN void* mod, IN uint32_t freq_div);
EXPORT_AUTD double AUTDModulationSamplingFrequency(IN const void* mod);
EXPORT_AUTD void AUTDDeleteModulation(IN const void* mod);
EXPORT_AUTD void AUTDPointSTM(OUT void** out);
EXPORT_AUTD void AUTDGainSTM(OUT void** out, IN const void* handle);
EXPORT_AUTD bool AUTDPointSTMAdd(IN void* stm, IN double x, IN double y, IN double z, IN uint8_t shift);
EXPORT_AUTD bool AUTDGainSTMAdd(IN void* stm, IN void* gain);
EXPORT_AUTD uint16_t AUTDGetGainSTMMode(IN void* stm);
EXPORT_AUTD void AUTDSetGainSTMMode(IN void* stm, IN uint16_t mode);
EXPORT_AUTD double AUTDSTMSetFrequency(IN void* stm, IN double freq);
EXPORT_AUTD double AUTDSTMFrequency(IN const void* stm);
EXPORT_AUTD double AUTDSTMSamplingFrequency(IN const void* stm);
EXPORT_AUTD uint32_t AUTDSTMSamplingFrequencyDivision(IN const void* stm);
EXPORT_AUTD void AUTDSTMSetSamplingFrequencyDivision(IN void* stm, IN uint32_t freq_div);
EXPORT_AUTD void AUTDDeleteSTM(IN const void* stm);
EXPORT_AUTD int32_t AUTDStop(IN void* handle);
EXPORT_AUTD void AUTDCreateSilencer(OUT void** out, IN uint16_t step, IN uint16_t cycle);
EXPORT_AUTD void AUTDDeleteSilencer(IN const void* config);
EXPORT_AUTD int32_t AUTDSend(IN void* handle, IN void* header, IN void* body);
EXPORT_AUTD void AUTDSetModDelay(IN void* handle, IN int32_t device_idx, IN int32_t local_trans_idx, IN uint16_t delay);
EXPORT_AUTD void AUTDCreateModDelayConfig(OUT void** out);
EXPORT_AUTD void AUTDDeleteModDelayConfig(IN const void* config);
EXPORT_AUTD void AUTDCreateAmplitudes(OUT void** out, IN void* handle, IN double amp);
EXPORT_AUTD void AUTDDeleteAmplitudes(IN const void* amplitudes);
EXPORT_AUTD void AUTDSetMode(IN uint8_t mode);
#ifdef __cplusplus
}
#endif
