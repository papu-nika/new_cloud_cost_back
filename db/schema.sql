
CREATE TYPE purchase_option as enum ('NoUpfront', 'PartialUpfront', 'AllUpfront');
CREATE TYPE offering_class as enum ('standard', 'convertible');

CREATE TABLE aws_ec2_inctances (
    id CHAR(16) PRIMARY KEY,
    -- productFamily VARCHAR(255) NOT NULL, // Compute Instance, Dedicated Host, Compute Instance (bare metal)
    regionCode VARCHAR(24) NOT NULL,
    instanceType VARCHAR(24) NOT NULL,
    instanceFamily VARCHAR(32) NOT NULL,
    vcpu INT NOT NULL,
    physicalProcessor VARCHAR(48) NOT NULL,
    clockSpeed VARCHAR(24) NOT NULL,
    memory real NOT NULL,
    storage VARCHAR(48) NOT NULL,
    networkPerformance VARCHAR(32) NOT NULL,
    operatingSystem VARCHAR(12) NOT NULL,
    preInstalledSw VARCHAR(8) NOT NULL,
    licenseModel VARCHAR(24) NOT NULL,
    capacitystatus VARCHAR(32) NOT NULL,
    tenancy VARCHAR(12) NOT NULL,
    dedicatedEbsThroughput VARCHAR(32) NOT NULL,
    ecu VARCHAR(8) NOT NULL,
    gpuMemory VARCHAR(24) NOT NULL,
    marketoption VARCHAR(24) NOT NULL, 
    processorFeatures VARCHAR(100) NOT NULL,
    ondemandPrice real NOT NULL,
    one_year_reserved_standard_price real NOT NULL,
    three_year_reserved_standard_price real NOT NULL,
    one_year_reserved_convertible_price real NOT NULL,
    three_year_reserved_convertible_price real NOT NULL
);

CREATE TABLE aws_ec2_inctance_reserved_prices  (
    id CHAR(16) PRIMARY KEY,
    LeaseContractLength int NOT NULL,
    purchaseOption purchase_option NOT NULL,
    offeringClass offering_class NOT NULL,
    reserved_price real
);

CREATE TABLE aws_ec2_nat_gateway (
    id CHAR(16) PRIMARY KEY,
    regionCode VARCHAR(24) NOT NULL,
    price real  
);

CREATE TABLE aws_rds_instances (
    id CHAR(16) PRIMARY KEY,
    regionCode VARCHAR(24) NOT NULL,
    instanceType VARCHAR(32) NOT NULL,
    instanceFamily VARCHAR(24) NOT NULL,
    vcpu INT NOT NULL,
    physicalProcessor VARCHAR(48) NOT NULL,
    clockSpeed VARCHAR(24) NOT NULL,
    memory real NOT NULL,
    storage VARCHAR(48) NOT NULL,
    networkPerformance VARCHAR(32) NOT NULL,
    databaseEngine VARCHAR(40) NOT NULL,
    licenseModel VARCHAR(32) NOT NULL,
    operation VARCHAR(32) NOT NULL,
    deploymentOption VARCHAR(32) NOT NULL,
    -- tenancy VARCHAR(32) NOT NULL,
    dedicatedEbsThroughput VARCHAR(32) NOT NULL,
    ondemandPrice real
);

CREATE TABLE aws_aurora_serverlesses (
    id CHAR(16) PRIMARY KEY,
    regionCode VARCHAR(24) NOT NULL,
    isAuroraIOOptimizationMode BOOLEAN NOT NULL,
    databaseEngine VARCHAR(40) NOT NULL,
    ondemandPrice real
);

CREATE TYPE lambda_type as enum ('duration', 'provisioned', 'edge-duration', 'edge-request',  'requests');


CREATE TABLE aws_lambdas (
    id CHAR(16) PRIMARY KEY,
    regionCode VARCHAR(24) NOT NULL,
    architecture VARCHAR(6) NOT NULL,
    type lambda_type NOT NULL,
    unit VARCHAR(16) NOT NULL,
    price real
);

CREATE TABLE sessions (
    id CHAR(24) PRIMARY KEY,
    user_id CHAR(32),
    data TEXT NOT NULL
    -- FOREIGN KEY (user_id) REFERENCES users(id)
);




    -- servicename CHAR : "Amazon Elastic Compute Cloud",
    -- location VARCHAR(24) NOT NULL,
    -- locationType CHAR : "AWS Region",
    -- currentGeneration CHAR : "Yes",
    -- tenancy CHAR : "Shared",
    -- licenseModel CHAR : "No License required",
    -- usagetype CHAR : "EUC1-BoxUsage:c5n.xlarge",
    -- availabilityzone CHAR : "NA",
    -- classicnetworkingsupport CHAR : "false",
    -- enhancedNetworkingSupported CHAR : "Yes",
    -- intelAvxAvailable CHAR : "Yes",
    -- intelAvx2Available CHAR : "Yes",
    -- intelTurboAvailable CHAR : "Yes",
    -- normalizationSizeFactor CHAR : "8",
    -- preInstalledSw CHAR : "NA",
    -- vpcetworkingsupport CHAR : "true"