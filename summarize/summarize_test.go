package summarize

import (
	"strings"
	"testing"
)

func TestSummarize(t *testing.T) {
	s := New(title, strings.NewReader(text))

	points := s.KeyPoints()
	checkKeyPoints(t, points, keyPoints)
}

func TestSummarizeFromString(t *testing.T) {
	s := NewFromString(title, text)

	points := s.KeyPoints()
	checkKeyPoints(t, points, keyPoints)

	dup := NewFromString(dupTitle, dupText)
	dupPoints := dup.KeyPoints()

	for _, p := range dupPoints {
		t.Log(p)
	}

	for i := 0; i < 10; i++ {
		dup := NewFromString(dupTitle, dupText)
		dupPoints := dup.KeyPoints()
		checkKeyPoints(t, dupPoints, dupKeyPoints)
	}
}

func checkKeyPoints(t *testing.T, points []string, kpoints []string) {
	if len(points) != len(kpoints) {
		t.Fatalf("Number of key points differ from expected: %d - %d\n", len(points), len(kpoints))
	}
	for i, k := range kpoints {
		if k != points[i] {
			t.Fatalf("Expected summary point '%s', got '%s'\n", k, points[i])
		}
	}
}

var (
	title = `Framework for Partitioning and Execution of Data Stream Applications in Mobile Cloud Computing`
	text  = `The contribution of cloud computing and mobile computing technologies lead to the newly emerging mobile cloud com- puting paradigm.
Three major approaches have been pro- posed for mobile cloud applications: 1) extending the access to cloud services to mobile devices; 2) enabling mobile de- vices to work collaboratively as cloud resource providers; 3) augmenting the execution of mobile applications on portable devices using cloud resources.
In this paper, we focus on the third approach in supporting mobile data stream applica- tions.
More specifically, we study how to optimize the com- putation partitioning of a data stream application between mobile and cloud to achieve maximum speed/throughput in processing the streaming data.
To the best of our knowledge, it is the first work to study the partitioning problem for mobile data stream applica- tions, where the optimization is placed on achieving high throughput of processing the streaming data rather than minimizing the makespan of executions as in other appli- cations.
We first propose a framework to provide runtime support for the dynamic computation partitioning and exe- cution of the application.
Different from existing works, the framework not only allows the dynamic partitioning for a single user but also supports the sharing of computation in- stances among multiple users in the cloud to achieve efficient utilization of the underlying cloud resources.
Meanwhile, the framework has better scalability because it is designed on the elastic cloud fabrics.
Based on the framework, we design a genetic algorithm for optimal computation parti- tion.
Both numerical evaluation and real world experiment have been performed, and the results show that the par- titioned application can achieve at least two times better performance in terms of throughput than the application without partitioning.`
	keyPoints = []string{
		`The contribution of cloud computing and mobile computing technologies lead to the newly emerging mobile cloud com- puting paradigm.`,
		`Three major approaches have been pro- posed for mobile cloud applications: 1) extending the access to cloud services to mobile devices; 2) enabling mobile de- vices to work collaboratively as cloud resource providers; 3) augmenting the execution of mobile applications on portable devices using cloud resources.`,
		`In this paper, we focus on the third approach in supporting mobile data stream applica- tions.`,
		`More specifically, we study how to optimize the com- putation partitioning of a data stream application between mobile and cloud to achieve maximum speed/throughput in processing the streaming data.`,
		`We first propose a framework to provide runtime support for the dynamic computation partitioning and exe- cution of the application.`,
	}
	dupTitle = "Cancer doc Farid Fata appeals 45-year prison sentence"
	dupText  = `A metro Detroit doctor who was sentenced to 45 years in federal prison last month is appealing his sentencing and conviction to the U.S. 6th Circuit Court of Appeals.
Farid Fata, 50, who was convicted of violating more than 550 patients' trust and raking in more than $17 million from fraudulent billings was sentenced on July 10.
Defense Attorney Mark Kriger filed the appeal on Fata's behalf  Wednesday in the U.S. District Court's Eastern District of Michigan. Kriger confirmed the filing to the Free Press.
"We filed a notice of appeal on the sentence, but I don't feel it's appropriate to comment on pending cases," Kriger said.
Prior to Fata's sentencing, the court heard victim impact statements from nearly 22 victims, who shared unfathomable experiences of undergoing unnecessary chemotherapy treatments and losing teeth, of a patient diagnosed with lung cancer when he had kidney cancer and more. One patient was given 195 chemotherapy treatments, 177 of which were unnecessary.
Fata pleaded guilty in September to 13 counts of health care fraud, two counts of money laundering and one count of conspiring to pay and receive kickbacks. The case involved $34.7 million in billings to patients and insurance companies, and $17.6 million paid for work Fata admitted was unnecessary.
Federal prosecutors said Fata's case was the most egregious fraud case they've ever seen. U.S. District Judge Paul Borman said before sentencing Fata that the once-prominent oncologist "practiced greed and shut down whatever compassion he had."
Fata, a married father of three and a naturalized U.S. citizen whose native country is Lebanon, was charged with running the scheme that involved billing the government for medically unnecessary cancer and blood treatments.
The government said Fata ran the scheme from 2009 to 2014 through his medical businesses, including Michigan Hematology Oncology Centers, with offices in Clarkston, Bloomfield Hills, Lapeer, Sterling Heights, Troy and Oak Park.
He remains incarcerated, and his medical license has been revoked, but Fata's legal troubles are far from over. About $13 million has been collected since 2013 to go toward a $17.6-million criminal judgement against Fata.
U.S. Assistant Prosecutor Catherine Dick said prosecutors are continuing to work to close the gap with his assets. Dick said the patient-victims and their families are first priority for compensation, then private insurers, then Medicare. The whistle-blower who tipped off the federal investigation is to receive 10% as part of an agreement; typically, whistle-blowers receive 15-25%, Dick said. Fata is also facing 27 pending lawsuits from patients and their families in Oakland County. George Karadsheh, Fata's former practice business manager, has also filed a whistle-blower federal suit against the former doctor.
Contact Katrease Stafford: kstafford@freepress.com
`
	dupKeyPoints = []string{
		`A metro Detroit doctor who was sentenced to 45 years in federal prison last month is appealing his sentencing and conviction to the U.`,
		`Farid Fata, 50, who was convicted of violating more than 550 patients' trust and raking in more than $17 million from fraudulent billings was sentenced on July 10.`,
		`Defense Attorney Mark Kriger filed the appeal on Fata's behalf  Wednesday in the U.`,
		`The government said Fata ran the scheme from 2009 to 2014 through his medical businesses, including Michigan Hematology Oncology Centers, with offices in Clarkston, Bloomfield Hills, Lapeer, Sterling Heights, Troy and Oak Park.`,
		`George Karadsheh, Fata's former practice business manager, has also filed a whistle-blower federal suit against the former doctor.`,
	}
)
