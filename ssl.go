package maestropanel

import "encoding/json"

type SSL struct {
	mp MaestroPanel
}

func (m *SSL) CreateSSLRequest(createSSL CreateSSL) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(createSSLRequestAction.Method, m.mp.getURL(createSSLRequestAction), createSSL)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *SSL) GetSSLCert(domainName string) (result SSLResult, err error) {
	result = SSLResult{}

	extra := struct {
		Name string `json:"name"`
	}{
		domainName,
	}

	response, err := m.mp.writeData(getSSLAction.Method, m.mp.getURL(getSSLAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}