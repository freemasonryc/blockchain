package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

type ProposalWrapper struct {
	Prop gov.Content
}

func TestContentAccessors(t *testing.T) {
	cases := map[string]struct {
		p     gov.Content
		title string
		desc  string
		typ   string
		str   string
	}{
		"upgrade": {
			p: types.NewSoftwareUpgradeProposal("Title", "desc", types.Plan{
				Name:   "due_height",
				Info:   "https:
				Height: 99999999999,
			}),
			title: "Title",
			desc:  "desc",
			typ:   "SoftwareUpgrade",
			str:   "Software Upgrade Proposal:\n  Title:       Title\n  Description: desc\n",
		},
		"cancel": {
			p:     types.NewCancelSoftwareUpgradeProposal("Cancel", "bad idea"),
			title: "Cancel",
			desc:  "bad idea",
			typ:   "CancelSoftwareUpgrade",
			str:   "Cancel Software Upgrade Proposal:\n  Title:       Cancel\n  Description: bad idea\n",
		},
	}

	cdc := codec.NewLegacyAmino()
	gov.RegisterLegacyAminoCodec(cdc)
	types.RegisterLegacyAminoCodec(cdc)

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.title, tc.p.GetTitle())
			assert.Equal(t, tc.desc, tc.p.GetDescription())
			assert.Equal(t, tc.typ, tc.p.ProposalType())
			assert.Equal(t, "upgrade", tc.p.ProposalRoute())
			assert.Equal(t, tc.str, tc.p.String())


			wrap := ProposalWrapper{tc.p}
			bz, err := cdc.Marshal(&wrap)
			require.NoError(t, err)
			unwrap := ProposalWrapper{}
			err = cdc.Unmarshal(bz, &unwrap)
			require.NoError(t, err)


			assert.Equal(t, tc.title, unwrap.Prop.GetTitle())
			assert.Equal(t, tc.desc, unwrap.Prop.GetDescription())
			assert.Equal(t, tc.typ, unwrap.Prop.ProposalType())
			assert.Equal(t, "upgrade", unwrap.Prop.ProposalRoute())
			assert.Equal(t, tc.str, unwrap.Prop.String())

		})

	}
}


func TestMarshalSoftwareUpdateProposal(t *testing.T) {

	plan := types.Plan{
		Name:   "upgrade",
		Height: 1000,
	}
	content := types.NewSoftwareUpgradeProposal("title", "description", plan)
	sup, ok := content.(*types.SoftwareUpgradeProposal)
	require.True(t, ok)


	ir := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(ir)
	gov.RegisterInterfaces(ir)
	cdc := codec.NewProtoCodec(ir)


	bz, err := cdc.MarshalJSON(sup)
	require.NoError(t, err)


	newSup := &types.SoftwareUpgradeProposal{}
	err = cdc.UnmarshalJSON(bz, newSup)
	require.NoError(t, err)
}
