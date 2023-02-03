/*
Wasp API

Testing CorecontractsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "./openapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_CorecontractsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CorecontractsApiService AccountsGetAccountBalance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var agentID string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetAccountBalance(context.Background(), chainID, agentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService AccountsGetAccountNFTIDs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var agentID string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetAccountNFTIDs(context.Background(), chainID, agentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService AccountsGetAccountNonce", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var agentID string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetAccountNonce(context.Background(), chainID, agentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService AccountsGetAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetAccounts(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService AccountsGetFoundryOutput", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var chainID2 string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetFoundryOutput(context.Background(), chainID, chainID2).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService AccountsGetNFTData", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var nftID string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetNFTData(context.Background(), chainID, nftID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService AccountsGetNativeTokenIDRegistry", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetNativeTokenIDRegistry(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService AccountsGetTotalAssets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.AccountsGetTotalAssets(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlobsGetAllBlobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlobsGetAllBlobs(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlobsGetBlobInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var blobHash string

		resp, httpRes, err := apiClient.CorecontractsApi.BlobsGetBlobInfo(context.Background(), chainID, blobHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlobsGetBlobValue", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var blobHash string
		var fieldKey string

		resp, httpRes, err := apiClient.CorecontractsApi.BlobsGetBlobValue(context.Background(), chainID, blobHash, fieldKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetBlockInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var blockIndex uint32

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetBlockInfo(context.Background(), chainID, blockIndex).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetControlAddresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetControlAddresses(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetEventsOfBlock", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var blockIndex uint32

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetEventsOfBlock(context.Background(), chainID, blockIndex).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetEventsOfContract", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var contractHname string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetEventsOfContract(context.Background(), chainID, contractHname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetEventsOfLatestBlock", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetEventsOfLatestBlock(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetEventsOfRequest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var requestID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetEventsOfRequest(context.Background(), chainID, requestID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetLatestBlockInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetLatestBlockInfo(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetRequestIDsForBlock", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var blockIndex uint32

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetRequestIDsForBlock(context.Background(), chainID, blockIndex).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetRequestIDsForLatestBlock", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetRequestIDsForLatestBlock(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetRequestIsProcessed", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var requestID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetRequestIsProcessed(context.Background(), chainID, requestID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetRequestReceipt", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var requestID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetRequestReceipt(context.Background(), chainID, requestID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetRequestReceiptsOfBlock", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var blockIndex uint32

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetRequestReceiptsOfBlock(context.Background(), chainID, blockIndex).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService BlocklogGetRequestReceiptsOfLatestBlock", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.BlocklogGetRequestReceiptsOfLatestBlock(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService ErrorsGetErrorMessageFormat", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string
		var contractHname string
		var errorID string

		resp, httpRes, err := apiClient.CorecontractsApi.ErrorsGetErrorMessageFormat(context.Background(), chainID, contractHname, errorID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorecontractsApiService GovernanceGetChainInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var chainID string

		resp, httpRes, err := apiClient.CorecontractsApi.GovernanceGetChainInfo(context.Background(), chainID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
