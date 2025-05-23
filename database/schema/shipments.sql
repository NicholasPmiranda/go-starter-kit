CREATE TABLE shipments (
    shipment_id VARCHAR NULL,
    trans VARCHAR NULL,
    cus_info VARCHAR NULL,
    mode VARCHAR NULL,
    origin VARCHAR NULL,
    origin_ctry VARCHAR NULL,
    destination VARCHAR NULL,
    dest_ctry VARCHAR NULL,
    consignor_code VARCHAR NULL,
    consignor_name VARCHAR NULL,
    consignee_code VARCHAR NULL,
    consignee_name VARCHAR NULL,
    house_ref VARCHAR NULL,
    inco VARCHAR NULL,
    additional_terms VARCHAR NULL,
    ppd_ccx VARCHAR NULL,
    goods_description VARCHAR NULL,
    origin_etd DATE NULL,
    dest_eta DATE NULL,
    weight NUMERIC NULL,
    weight_uq VARCHAR NULL,
    volume NUMERIC NULL,
    volume_uq VARCHAR NULL,
    loading_meters NUMERIC NULL,
    chargeable NUMERIC NULL,
    chargeable_uq VARCHAR NULL,
    added DATE NULL,
    controlling_customer_code VARCHAR NULL,
    controlling_customer_name VARCHAR NULL,
    controlling_agent_code VARCHAR NULL,
    controlling_agent_name VARCHAR NULL,
    transport_job VARCHAR NULL,
    brokerage_job VARCHAR NULL,
    is_master_lead BOOLEAN NULL,
    master_lead_ref VARCHAR NULL,
    import_broker_code VARCHAR NULL,
    import_broker_name VARCHAR NULL,
    export_broker_code VARCHAR NULL,
    export_broker_name VARCHAR NULL,
    job_branch VARCHAR NULL,
    job_dept VARCHAR NULL,
    local_client_code VARCHAR NULL,
    local_client_name VARCHAR NULL,
    job_sales_rep VARCHAR NULL,
    job_operator VARCHAR NULL,
    job_status VARCHAR NULL,
    job_opened DATE NULL,
    recognized_revenue NUMERIC NULL,
    recognized_wip NUMERIC NULL,
    total_recognized_income NUMERIC NULL,
    recognized_cost NUMERIC NULL,
    recognized_accrual NUMERIC NULL,
    total_recognized_expense NUMERIC NULL,
    job_profit NUMERIC NULL,
    consol_id VARCHAR NULL,
    first_load VARCHAR NULL,
    last_disch VARCHAR NULL,
    etd_first_load DATE NULL,
    eta_last_disch DATE NULL,
    master VARCHAR NULL,
    vessel VARCHAR NULL,
    flight_voyage VARCHAR NULL,
    load VARCHAR NULL,
    disch VARCHAR NULL,
    etd_load DATE NULL,
    eta_disch DATE NULL,
    send_agent_code VARCHAR NULL,
    send_agent_name VARCHAR NULL,
    recv_agent VARCHAR NULL,
    recv_agent_name VARCHAR NULL,
    co_loaded_with VARCHAR NULL,
    co_loader_name VARCHAR NULL,
    carrier_code VARCHAR NULL,
    carrier_name VARCHAR NULL,
    teu NUMERIC NULL,
    cntr_count NUMERIC NULL,
    other NUMERIC NULL,
    f20 NUMERIC NULL,
    r20 NUMERIC NULL,
    h20 NUMERIC NULL,
    f40 NUMERIC NULL,
    r40 NUMERIC NULL,
    h40 NUMERIC NULL,
    f45 NUMERIC NULL,
    gen NUMERIC NULL,
    unrecognized_revenue NUMERIC NULL,
    unrecognized_wip NUMERIC NULL,
    unrecognized_cost NUMERIC NULL,
    unrecognized_accrual NUMERIC NULL,
    total_revenue NUMERIC NULL,
    total_wip NUMERIC NULL,
    total_income NUMERIC NULL,
    service_level_code VARCHAR NULL,
    shippers_reference VARCHAR NULL,
    consignor_city VARCHAR NULL,
    consignor_state VARCHAR NULL,
    consignor_postcode VARCHAR NULL,
    consignee_city VARCHAR NULL,
    consignee_state VARCHAR NULL,
    consignee_postcode VARCHAR NULL,
    consol_atd DATE NULL,
    consol_ata DATE NULL,
    job_revenue_recognition_date VARCHAR NULL,
    direction VARCHAR NULL,
    local_client_ar_settlement_group_code VARCHAR NULL,
    local_client_ar_settlement_group_name VARCHAR NULL,
    overseas_agent_code VARCHAR NULL,
    overseas_agent_name VARCHAR NULL,
    job_overseas_agent_ar_settlement_group_code VARCHAR NULL,
    job_overseas_agent_ar_settlement_group_name VARCHAR NULL,
    total_cost NUMERIC NULL,
    total_accrual NUMERIC NULL,
    total_expense NUMERIC NULL
);

CREATE INDEX idx_shipment_id ON shipments (shipment_id);
