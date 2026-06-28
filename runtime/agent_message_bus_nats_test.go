package agentruntime

import (
	"testing"

	"github.com/nats-io/nats.go/jetstream"
)

func TestAgentMessageConsumerConfigUsesNewDeliverPolicy(t *testing.T) {
	cfg := agentMessageConsumerConfig(
		"agentflow-worker-a-konstrukt-system-analysis-konstrukt-document-intake-agent",
		"orloj.agentmsg.konstrukt-system-analysis.konstrukt-document-intake-agent.inbox",
	)

	if cfg.DeliverPolicy != jetstream.DeliverNewPolicy {
		t.Fatalf("expected DeliverNewPolicy, got %v", cfg.DeliverPolicy)
	}
	if cfg.AckPolicy != jetstream.AckExplicitPolicy {
		t.Fatalf("expected AckExplicitPolicy, got %v", cfg.AckPolicy)
	}
	if cfg.AckWait != defaultConsumerAckWait {
		t.Fatalf("expected AckWait %s, got %s", defaultConsumerAckWait, cfg.AckWait)
	}
}

func TestDurableNameUsesCurrentConfigVersion(t *testing.T) {
	got := durableName("agentflow-worker-a", "konstrukt-system-analysis", "konstrukt-document-intake-agent")
	want := "agentflow-worker-a-konstrukt-system-analysis-konstrukt-document-intake-agent-v2"
	if got != want {
		t.Fatalf("expected durable %q, got %q", want, got)
	}
}
