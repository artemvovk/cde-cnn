package types

type DataElement struct {
	VariableName                                  string `csv:"variable name"`
	Title                                         string `csv:"title"`
	ElementType                                   string `csv:"element type"`
	Version                                       string `csv:"version"`
	Definition                                    string `csv:"definition"`
	ShortDescription                              string `csv:"short description"`
	Datatype                                      string `csv:"datatype"`
	MaximumCharacterQuantity                      string `csv:"maximum character quantity"`
	InputRestriction                              string `csv:"input restriction"`
	MinimumValue                                  string `csv:"minimum value"`
	MaximumValue                                  string `csv:"maximum value"`
	PermissibleValues                             string `csv:"permissible values"`
	PermissibleValueDescriptions                  string `csv:"permissible value descriptions"`
	PermissibleValueOutputCodes                   string `csv:"permissible value output codes"`
	ItemResponseOID                               string `csv:"item response OID"`
	ElementOID                                    string `csv:"element OID"`
	UnitOfMeasure                                 string `csv:"unite of measure"`
	GuidelinesInstructions                        string `csv:"guidelines"`
	Notes                                         string `csv:"notes"`
	PreferredQuestionText                         string `csv:"preferred question text"`
	Keywords                                      string `csv:"keywords"`
	References                                    string `csv:"references"`
	PopulationAll                                 string `csv:"poplation.all"`
	DomainGeneral                                 string `csv:"domain.general"`
	DomainTraumaticBrainInjury                    string `csv:"domain.traumatic brain injury"`
	DomainParkinsonsDisease                       string `csv:"domain.pakrinsons disease"`
	DomainFriedreichsAtaxia                       string `csv:"domain.friedirchs ataxia"`
	DomainStroke                                  string `csv:"domain.stroke"`
	DomainAmyotrophicLateralSclerosis             string `csv:"domain.amyotrophc lateral sclerosis"`
	DomainHuntingtonsDisease                      string `csv:"domain.huntingtons disease"`
	DomainMultipleSclerosis                       string `csv:"domain.multiple sclerosis"`
	DomainNeuromuscularDiseases                   string `csv:"domain.neuromuscular diseases"`
	DomainMyastheniaGravis                        string `csv:"domain.myasthenia gravis"`
	DomainSpinalMuscularAtrophy                   string `csv:"domain.spinal muscular atrophy"`
	DomainDuchenneBeckerMuscularDystrophy         string `csv:"domain.duchenne becker muscular dystrophy"`
	DomainCongenitalMuscularDystrophy             string `csv:"domain.congenital muscular dystrophy"`
	DomainSpinalCordInjury                        string `csv:"domain.spinal cord injury"`
	DomainHeadache                                string `csv:"domain.headache"`
	DomainEpilepsy                                string `csv:"domain.epilepsy"`
	ClassificationGeneral                         string `csv:"classification.general"`
	ClassificationAcuteHospitalized               string `csv:"classification.acute hospitalized"`
	ClassificationConcussionMildTBI               string `csv:"classification.concussion mild TBI"`
	ClassificationEpidemiology                    string `csv:"classification.epidemiology"`
	ClassificationModerateSevereTBIRehab          string `csv:"classification.moderate severe TBI rehab"`
	ClassificationParkinsonsDisease               string `csv:"classification.parkinsons-disease"`
	ClassificationFriedreichsAtaxia               string `csv:"classification.friedreichs ataxia"`
	ClassificationStroke                          string `csv:"classification.stroke"`
	ClassificationAmyotrophicLateralSclerosis     string `csv:"classification.amyotrophic lateral sclerosis"`
	ClassificationHuntingtonsDisease              string `csv:"classification.huntingtons disease"`
	ClassificationMultipleSclerosis               string `csv:"classification.multiple sclerosis"`
	ClassificationNeuromuscularDiseases           string `csv:"classification.neuromuscular diseases"`
	ClassificationMyastheniaGravis                string `csv:"classification.myasthenia gravis"`
	ClassificationSpinalMuscularAtrophy           string `csv:"classification.spinal muscular atrophy"`
	ClassificationDuchenneBeckerMuscularDystrophy string `csv:"classification.duchenne becker muscular dystrophy"`
	ClassificationCongenitalMuscularDystrophy     string `csv:"classification.congenital muscular dystrphy"`
	ClassificationSpinalCordInjury                string `csv:"classification.spinal cord injury"`
	ClassificationHeadache                        string `csv:"classification.headache"`
	ClassificationEpilepsy                        string `csv:"classification.epilepsy"`
	HistoricalNotes                               string `csv:"historical notes"`
	Labels                                        string `csv:"labels"`
	SeeAlso                                       string `csv:"see also"`
	SubmittingOrganizationName                    string `csv:"submitting organization name"`
	SubmittingContactName                         string `csv:"submitting contact name"`
	SubmittingContactInformation                  string `csv:"submitting contact information"`
	EffectiveDate                                 string `csv:"effective date"`
	UntilDate                                     string `csv:"until date"`
	StewardOrganizationName                       string `csv:"steward organiation name"`
	StewardContactName                            string `csv:"steward contact name"`
	StewardContactInformation                     string `csv:"steward contact information"`
	CreationDate                                  string `csv:"creation  date"`
	LastChangeDate                                string `csv:"last change date"`
	AdministrativeStatus                          string `csv:"administrative status"`
	CATandOID                                     string `csv:"CAT OID"`
	FormItemOID                                   string `csv:"form item OID"`
}

type Token struct {
	Text  string
	Tag   string
	Label string
}

type TermFrequency struct {
	Word      Token
	Frequency float64
}

type DocFrequency struct {
	Doc       DataElement
	Frequency float64
}
