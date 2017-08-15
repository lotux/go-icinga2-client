package icinga2

type CheckResult struct {
    ExitStatus int `json:"exit_status"`
    PluginOutput string `json:"plugin_output"`
    PerformanceData []string `json:"performance_data"`
    CheckSource string `json:"check_source"`
}

func (s *WebClient) ProcessCheckResult(host string, service string, checkResult CheckResult) error {
    serviceName := checkResult.CheckSource + "!" + service
    err := s.DoAction("/process-check-result?service="+serviceName, checkResult)
    if err != nil {
        return err
    }

    return nil
}